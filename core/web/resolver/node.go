package resolver

import (
	"context"

	"github.com/graph-gophers/graphql-go"

	"github.com/smartcontractkit/chainlink/core/chains/evm/types"
	"github.com/smartcontractkit/chainlink/core/web/loader"
)

// NodeResolver resolves the Node type.
type NodeResolver struct {
	node types.Node
}

func NewNode(node types.Node) *NodeResolver {
	return &NodeResolver{node: node}
}

func NewNodes(nodes []types.Node) []*NodeResolver {
	resolvers := []*NodeResolver{}
	for _, n := range nodes {
		resolvers = append(resolvers, NewNode(n))
	}

	return resolvers
}

// ID resolves the node's unique identifier.
func (r *NodeResolver) ID() graphql.ID {
	return int32GQLID(r.node.ID)
}

// Name resolves the node's name field.
func (r *NodeResolver) Name() string {
	return r.node.Name
}

// WSURL resolves the node's websocket url field.
func (r *NodeResolver) WSURL() string {
	return r.node.WSURL.String
}

// HTTPURL resolves the node's http url field.
func (r *NodeResolver) HTTPURL() string {
	return r.node.HTTPURL.String
}

// Chain resolves the node's chain object field.
func (r *NodeResolver) Chain(ctx context.Context) (*ChainResolver, error) {
	chain, err := loader.GetChainByID(ctx, r.node.EVMChainID.String())
	if err != nil {
		return nil, err
	}

	return NewChain(*chain), nil
}

// CreatedAt resolves the node's created at field.
func (r *NodeResolver) CreatedAt() graphql.Time {
	return graphql.Time{Time: r.node.CreatedAt}
}

// UpdatedAt resolves the node's updated at field.
func (r *NodeResolver) UpdatedAt() graphql.Time {
	return graphql.Time{Time: r.node.UpdatedAt}
}

// -- Node Query --

type NodePayloadResolver struct {
	node *types.Node
	err  error
}

func NewNodePayloadResolver(node *types.Node, err error) *NodePayloadResolver {
	return &NodePayloadResolver{node, err}
}

// ToNode resolves the Node object to be returned if it is found
func (r *NodePayloadResolver) ToNode() (*NodeResolver, bool) {
	if r.node != nil {
		return NewNode(*r.node), true
	}

	return nil, false
}

// ToNotFoundError implements the NotFoundError union type of the payload
func (r *NodePayloadResolver) ToNotFoundError() (*NotFoundErrorResolver, bool) {
	if r.err != nil {
		return NewNotFoundError("node not found"), true
	}

	return nil, false
}

// -- CreateNode Mutation --

type CreateNodePayloadResolver struct {
	node *types.Node
}

func NewCreateNodePayloadResolver(node *types.Node) *CreateNodePayloadResolver {
	return &CreateNodePayloadResolver{node}
}

func (r *CreateNodePayloadResolver) ToCreateNodeSuccess() (*CreateNodeSuccessResolve, bool) {
	if r.node != nil {
		return NewCreateNodeSuccessResolve(*r.node), true
	}

	return nil, false
}

type CreateNodeSuccessResolve struct {
	node types.Node
}

func NewCreateNodeSuccessResolve(node types.Node) *CreateNodeSuccessResolve {
	return &CreateNodeSuccessResolve{node}
}

func (r *CreateNodeSuccessResolve) Node() *NodeResolver {
	return NewNode(r.node)
}

// -- DeleteNode Mutation --

type DeleteNodePayloadResolver struct {
	node *types.Node
	err  error
}

func NewDeleteNodePayloadResolver(node *types.Node, err error) *DeleteNodePayloadResolver {
	return &DeleteNodePayloadResolver{node, err}
}

func (r *DeleteNodePayloadResolver) ToDeleteNodeSuccess() (*DeleteNodeSuccessResolver, bool) {
	if r.node != nil {
		return NewDeleteNodeSuccessResolver(r.node), true
	}

	return nil, false
}

func (r *DeleteNodePayloadResolver) ToNotFoundError() (*NotFoundErrorResolver, bool) {
	if r.err != nil {
		return NewNotFoundError("node not found"), true
	}

	return nil, false
}

type DeleteNodeSuccessResolver struct {
	node *types.Node
}

func NewDeleteNodeSuccessResolver(node *types.Node) *DeleteNodeSuccessResolver {
	return &DeleteNodeSuccessResolver{node}
}

func (r *DeleteNodeSuccessResolver) Node() *NodeResolver {
	return NewNode(*r.node)
}
