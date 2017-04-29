// *** WARNING: this file was generated by the Coconut IDL Compiler (CIDLC).  ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
    "errors"

    pbempty "github.com/golang/protobuf/ptypes/empty"
    pbstruct "github.com/golang/protobuf/ptypes/struct"
    "golang.org/x/net/context"

    "github.com/pulumi/coconut/pkg/resource"
    "github.com/pulumi/coconut/pkg/tokens"
    "github.com/pulumi/coconut/pkg/util/contract"
    "github.com/pulumi/coconut/pkg/util/mapper"
    "github.com/pulumi/coconut/sdk/go/pkg/cocorpc"
)

/* RPC stubs for BasePathMapping resource provider */

// BasePathMappingToken is the type token corresponding to the BasePathMapping package type.
const BasePathMappingToken = tokens.Type("aws:apigateway/basePathMapping:BasePathMapping")

// BasePathMappingProviderOps is a pluggable interface for BasePathMapping-related management functionality.
type BasePathMappingProviderOps interface {
    Check(ctx context.Context, obj *BasePathMapping) ([]mapper.FieldError, error)
    Create(ctx context.Context, obj *BasePathMapping) (string, error)
    Get(ctx context.Context, id string) (*BasePathMapping, error)
    InspectChange(ctx context.Context,
        id string, old *BasePathMapping, new *BasePathMapping, diff *resource.ObjectDiff) ([]string, error)
    Update(ctx context.Context,
        id string, old *BasePathMapping, new *BasePathMapping, diff *resource.ObjectDiff) error
    Delete(ctx context.Context, id string) error
}

// BasePathMappingProvider is a dynamic gRPC-based plugin for managing BasePathMapping resources.
type BasePathMappingProvider struct {
    ops BasePathMappingProviderOps
}

// NewBasePathMappingProvider allocates a resource provider that delegates to a ops instance.
func NewBasePathMappingProvider(ops BasePathMappingProviderOps) cocorpc.ResourceProviderServer {
    contract.Assert(ops != nil)
    return &BasePathMappingProvider{ops: ops}
}

func (p *BasePathMappingProvider) Check(
    ctx context.Context, req *cocorpc.CheckRequest) (*cocorpc.CheckResponse, error) {
    contract.Assert(req.GetType() == string(BasePathMappingToken))
    obj, _, decerr := p.Unmarshal(req.GetProperties())
    if decerr == nil || len(decerr.Failures()) == 0 {
        failures, err := p.ops.Check(ctx, obj)
        if err != nil {
            return nil, err
        }
        if len(failures) > 0 {
            decerr = mapper.NewDecodeErr(failures)
        }
    }
    return resource.NewCheckResponse(decerr), nil
}

func (p *BasePathMappingProvider) Name(
    ctx context.Context, req *cocorpc.NameRequest) (*cocorpc.NameResponse, error) {
    contract.Assert(req.GetType() == string(BasePathMappingToken))
    obj, _, decerr := p.Unmarshal(req.GetProperties())
    if decerr != nil {
        return nil, decerr
    }
    if obj.Name == "" {
        return nil, errors.New("Name property cannot be empty")
    }
    return &cocorpc.NameResponse{Name: obj.Name}, nil
}

func (p *BasePathMappingProvider) Create(
    ctx context.Context, req *cocorpc.CreateRequest) (*cocorpc.CreateResponse, error) {
    contract.Assert(req.GetType() == string(BasePathMappingToken))
    obj, _, decerr := p.Unmarshal(req.GetProperties())
    if decerr != nil {
        return nil, decerr
    }
    id, err := p.ops.Create(ctx, obj)
    if err != nil {
        return nil, err
    }
    return &cocorpc.CreateResponse{
        Id:   id,
    }, nil
}

func (p *BasePathMappingProvider) Get(
    ctx context.Context, req *cocorpc.GetRequest) (*cocorpc.GetResponse, error) {
    contract.Assert(req.GetType() == string(BasePathMappingToken))
    id := req.GetId()
    obj, err := p.ops.Get(ctx, id)
    if err != nil {
        return nil, err
    }
    return &cocorpc.GetResponse{
        Properties: resource.MarshalProperties(
            nil, resource.NewPropertyMap(obj), resource.MarshalOptions{}),
    }, nil
}

func (p *BasePathMappingProvider) InspectChange(
    ctx context.Context, req *cocorpc.ChangeRequest) (*cocorpc.InspectChangeResponse, error) {
    contract.Assert(req.GetType() == string(BasePathMappingToken))
    old, oldprops, decerr := p.Unmarshal(req.GetOlds())
    if decerr != nil {
        return nil, decerr
    }
    new, newprops, decerr := p.Unmarshal(req.GetNews())
    if decerr != nil {
        return nil, decerr
    }
    diff := oldprops.Diff(newprops)
    var replaces []string
    if diff.Changed("name") {
        replaces = append(replaces, "name")
    }
    more, err := p.ops.InspectChange(ctx, req.GetId(), old, new, diff)
    if err != nil {
        return nil, err
    }
    return &cocorpc.InspectChangeResponse{
        Replaces: append(replaces, more...),
    }, err
}

func (p *BasePathMappingProvider) Update(
    ctx context.Context, req *cocorpc.ChangeRequest) (*pbempty.Empty, error) {
    contract.Assert(req.GetType() == string(BasePathMappingToken))
    old, oldprops, err := p.Unmarshal(req.GetOlds())
    if err != nil {
        return nil, err
    }
    new, newprops, err := p.Unmarshal(req.GetNews())
    if err != nil {
        return nil, err
    }
    diff := oldprops.Diff(newprops)
    if err := p.ops.Update(ctx, req.GetId(), old, new, diff); err != nil {
        return nil, err
    }
    return &pbempty.Empty{}, nil
}

func (p *BasePathMappingProvider) Delete(
    ctx context.Context, req *cocorpc.DeleteRequest) (*pbempty.Empty, error) {
    contract.Assert(req.GetType() == string(BasePathMappingToken))
    if err := p.ops.Delete(ctx, req.GetId()); err != nil {
        return nil, err
    }
    return &pbempty.Empty{}, nil
}

func (p *BasePathMappingProvider) Unmarshal(
    v *pbstruct.Struct) (*BasePathMapping, resource.PropertyMap, mapper.DecodeError) {
    var obj BasePathMapping
    props := resource.UnmarshalProperties(v)
    result := mapper.MapIU(props.Mappable(), &obj)
    return &obj, props, result
}

/* Marshalable BasePathMapping structure(s) */

// BasePathMapping is a marshalable representation of its corresponding IDL type.
type BasePathMapping struct {
    Name string `json:"name"`
    DomainName string `json:"domainName"`
    RestAPI *resource.ID `json:"restAPI"`
    BasePath *string `json:"basePath,omitempty"`
    Stage *resource.ID `json:"stage,omitempty"`
}

// BasePathMapping's properties have constants to make dealing with diffs and property bags easier.
const (
    BasePathMapping_Name = "name"
    BasePathMapping_DomainName = "domainName"
    BasePathMapping_RestAPI = "restAPI"
    BasePathMapping_BasePath = "basePath"
    BasePathMapping_Stage = "stage"
)

