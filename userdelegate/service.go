package userdelegate

import (
	"context"
	"github.com/airbloc/airbloc-go/account"
	"github.com/airbloc/airbloc-go/apps"
	"github.com/airbloc/airbloc-go/collections"
	ablCommon "github.com/airbloc/airbloc-go/common"
	"github.com/airbloc/airbloc-go/dauth"
	"github.com/airbloc/airbloc-go/node"
	"github.com/airbloc/airbloc-go/p2p"
	p2pcommon "github.com/airbloc/airbloc-go/p2p/common"
	pb "github.com/airbloc/airbloc-go/proto/p2p/v1"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
	"log"
)

var (
	ErrDelegationNotAllowed = errors.New("the user is not designated you as a delegate.")
)

type Service struct {
	accountIds []ablCommon.ID
	p2p        p2p.Server
	selfAddr   ethCommon.Address
	isRunning  bool

	// managers for blockchain interaction
	apps        *apps.Manager
	dauth       *dauth.Manager
	accounts    *account.Manager
	collections *collections.Collections
}

func NewService(backend node.Backend) (node.Service, error) {
	return &Service{
		accountIds:  []ablCommon.ID{},
		p2p:         backend.P2P(),
		selfAddr:    backend.Kms().NodeKey().EthereumAddress,
		apps:        apps.NewManager(backend.Client()),
		dauth:       dauth.NewManager(backend.Client()),
		accounts:    account.NewManager(backend.Client()),
		collections: collections.New(backend.LocalDatabase(), backend.MetaDatabase(), backend.Client()),
	}, nil
}

// AddUser adds a user to the delegated user list,
// therefore manage
func (service *Service) AddUser(accId string) error {
	accountId, err := ablCommon.IDFromString(accId)
	if err != nil {
		return errors.Wrapf(err, "invalid account ID %s", accId)
	}

	// you can delegate a user after the user setting you as a delegate.
	if isDelegate, err := service.accounts.IsDelegateOf(service.selfAddr, accountId); err != nil {
		return errors.Wrapf(err, "failed to call Accounts.IsDelegateOf")
	} else if !isDelegate {
		return ErrDelegationNotAllowed
	}

	service.accountIds = append(service.accountIds, accountId)
	if service.isRunning {
		service.registerDAuthHandler(accountId)
	}
	return nil
}

func (service *Service) Start() error {
	service.p2p.SubscribeTopic("dauth-signup", &pb.DAuthSignUpRequest{}, service.signUpHandler)

	for _, accountId := range service.accountIds {
		service.registerDAuthHandler(accountId)
	}
}

func (service *Service) registerDAuthHandler(accountId ablCommon.ID) {
	accId := accountId.String()
	dauthReq := &pb.DAuthRequest{}

	service.p2p.SubscribeTopic("dauth-allow-" + accId, dauthReq, service.createDAuthHandler(accountId, true))
	service.p2p.SubscribeTopic("dauth-deny-" + accId, dauthReq, service.createDAuthHandler(accountId, false))
}

func (service *Service) createDAuthHandler(accountId ablCommon.ID, allow bool) p2p.TopicHandler {
	return func(server p2p.Server, ctx context.Context, message p2pcommon.Message) {
		request, ok := message.Data.(*pb.DAuthRequest)
		if !ok {
			log.Println("error: Invalid topic.")
			return
		}

		collectionId, err := ablCommon.IDFromString(request.CollectionId)
		if err != nil {
			log.Println("error: Invalid Collection ID", collectionId, err.Error())
			return
		}

		// the message sender should be the data provider (the collection's owner)
		senderAddr := crypto.PubkeyToAddress(*message.Sender)
		if ok, err := service.isCollectionOwner(ctx, collectionId, senderAddr); err != nil {
			log.Println("error: Failed to retrieve collection owner", err.Error())
			return
		} else if !ok {
			log.Println("error: The address", senderAddr.Hex(), "is not a data provider.")
			return
		}

		if allow {
			err = service.dauth.AllowByDelegate(collectionId, accountId)
		} else {
			err = service.dauth.DenyByDelegate(collectionId, accountId)
		}
		if err != nil {
			log.Println("error: Failed to modify DAuth settings: ", err.Error())
		}

		// TODO: reply to the data provider :D
	}
}

func (service *Service) signUpHandler(server p2p.Server, ctx context.Context, message p2pcommon.Message) {
	request, ok := message.Data.(*pb.DAuthSignUpRequest)
	if !ok {
		log.Println("error: Invalid topic.")
		return
	}

	identityHash := ethCommon.HexToHash(request.GetIdentityHash())
	accountId, err := service.accounts.CreateTemporary(identityHash)
	if err != nil {
		log.Println("error: Failed to create temporary account:", err.Error())
	}

	log.Println(
		"Created account", accountId.String(),
		"by request from the data provider", crypto.PubkeyToAddress(*message.Sender).Hex())

	response := &pb.DAuthSignUpResponse{
		UserId: accountId.String(),
	}
	if err = server.Send(context.Background(), response, "dauth-signup-response", message.Info.ID); err != nil {
		log.Println("error: Failed to send response to data provider:", err.Error())
	}
}

// isCollectionOwner checks that the P2P message sender is
// same with the owner of the collection (data provider, app owner).
func (service *Service) isCollectionOwner(ctx context.Context, collectionId ablCommon.ID, senderAddr ethCommon.Address) (bool, error) {
	collection, err := service.collections.Get(collectionId)
	if err != nil {
		return false, errors.Wrap(err, "unable to retrieve collection")
	}
	return service.apps.CheckOwner(ctx, collection.AppId, senderAddr)
}

func (service *Service) Stop() {
}
