// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// JungleTVClient is the client API for JungleTV service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type JungleTVClient interface {
	SignIn(ctx context.Context, in *SignInRequest, opts ...grpc.CallOption) (JungleTV_SignInClient, error)
	EnqueueMedia(ctx context.Context, in *EnqueueMediaRequest, opts ...grpc.CallOption) (*EnqueueMediaResponse, error)
	MonitorTicket(ctx context.Context, in *MonitorTicketRequest, opts ...grpc.CallOption) (JungleTV_MonitorTicketClient, error)
	ConsumeMedia(ctx context.Context, in *ConsumeMediaRequest, opts ...grpc.CallOption) (JungleTV_ConsumeMediaClient, error)
	MonitorQueue(ctx context.Context, in *MonitorQueueRequest, opts ...grpc.CallOption) (JungleTV_MonitorQueueClient, error)
	RewardInfo(ctx context.Context, in *RewardInfoRequest, opts ...grpc.CallOption) (*RewardInfoResponse, error)
	SubmitActivityChallenge(ctx context.Context, in *SubmitActivityChallengeRequest, opts ...grpc.CallOption) (*SubmitActivityChallengeResponse, error)
	ConsumeChat(ctx context.Context, in *ConsumeChatRequest, opts ...grpc.CallOption) (JungleTV_ConsumeChatClient, error)
	SendChatMessage(ctx context.Context, in *SendChatMessageRequest, opts ...grpc.CallOption) (*SendChatMessageResponse, error)
	SubmitProofOfWork(ctx context.Context, in *SubmitProofOfWorkRequest, opts ...grpc.CallOption) (*SubmitProofOfWorkResponse, error)
	// moderation endpoints
	ForciblyEnqueueTicket(ctx context.Context, in *ForciblyEnqueueTicketRequest, opts ...grpc.CallOption) (*ForciblyEnqueueTicketResponse, error)
	RemoveQueueEntry(ctx context.Context, in *RemoveQueueEntryRequest, opts ...grpc.CallOption) (*RemoveQueueEntryResponse, error)
	RemoveChatMessage(ctx context.Context, in *RemoveChatMessageRequest, opts ...grpc.CallOption) (*RemoveChatMessageResponse, error)
	SetChatSettings(ctx context.Context, in *SetChatSettingsRequest, opts ...grpc.CallOption) (*SetChatSettingsResponse, error)
	SetVideoEnqueuingEnabled(ctx context.Context, in *SetVideoEnqueuingEnabledRequest, opts ...grpc.CallOption) (*SetVideoEnqueuingEnabledResponse, error)
	BanUser(ctx context.Context, in *BanUserRequest, opts ...grpc.CallOption) (*BanUserResponse, error)
	RemoveBan(ctx context.Context, in *RemoveBanRequest, opts ...grpc.CallOption) (*RemoveBanResponse, error)
}

type jungleTVClient struct {
	cc grpc.ClientConnInterface
}

func NewJungleTVClient(cc grpc.ClientConnInterface) JungleTVClient {
	return &jungleTVClient{cc}
}

func (c *jungleTVClient) SignIn(ctx context.Context, in *SignInRequest, opts ...grpc.CallOption) (JungleTV_SignInClient, error) {
	stream, err := c.cc.NewStream(ctx, &_JungleTV_serviceDesc.Streams[0], "/jungletv.JungleTV/SignIn", opts...)
	if err != nil {
		return nil, err
	}
	x := &jungleTVSignInClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type JungleTV_SignInClient interface {
	Recv() (*SignInProgress, error)
	grpc.ClientStream
}

type jungleTVSignInClient struct {
	grpc.ClientStream
}

func (x *jungleTVSignInClient) Recv() (*SignInProgress, error) {
	m := new(SignInProgress)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *jungleTVClient) EnqueueMedia(ctx context.Context, in *EnqueueMediaRequest, opts ...grpc.CallOption) (*EnqueueMediaResponse, error) {
	out := new(EnqueueMediaResponse)
	err := c.cc.Invoke(ctx, "/jungletv.JungleTV/EnqueueMedia", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jungleTVClient) MonitorTicket(ctx context.Context, in *MonitorTicketRequest, opts ...grpc.CallOption) (JungleTV_MonitorTicketClient, error) {
	stream, err := c.cc.NewStream(ctx, &_JungleTV_serviceDesc.Streams[1], "/jungletv.JungleTV/MonitorTicket", opts...)
	if err != nil {
		return nil, err
	}
	x := &jungleTVMonitorTicketClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type JungleTV_MonitorTicketClient interface {
	Recv() (*EnqueueMediaTicket, error)
	grpc.ClientStream
}

type jungleTVMonitorTicketClient struct {
	grpc.ClientStream
}

func (x *jungleTVMonitorTicketClient) Recv() (*EnqueueMediaTicket, error) {
	m := new(EnqueueMediaTicket)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *jungleTVClient) ConsumeMedia(ctx context.Context, in *ConsumeMediaRequest, opts ...grpc.CallOption) (JungleTV_ConsumeMediaClient, error) {
	stream, err := c.cc.NewStream(ctx, &_JungleTV_serviceDesc.Streams[2], "/jungletv.JungleTV/ConsumeMedia", opts...)
	if err != nil {
		return nil, err
	}
	x := &jungleTVConsumeMediaClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type JungleTV_ConsumeMediaClient interface {
	Recv() (*MediaConsumptionCheckpoint, error)
	grpc.ClientStream
}

type jungleTVConsumeMediaClient struct {
	grpc.ClientStream
}

func (x *jungleTVConsumeMediaClient) Recv() (*MediaConsumptionCheckpoint, error) {
	m := new(MediaConsumptionCheckpoint)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *jungleTVClient) MonitorQueue(ctx context.Context, in *MonitorQueueRequest, opts ...grpc.CallOption) (JungleTV_MonitorQueueClient, error) {
	stream, err := c.cc.NewStream(ctx, &_JungleTV_serviceDesc.Streams[3], "/jungletv.JungleTV/MonitorQueue", opts...)
	if err != nil {
		return nil, err
	}
	x := &jungleTVMonitorQueueClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type JungleTV_MonitorQueueClient interface {
	Recv() (*Queue, error)
	grpc.ClientStream
}

type jungleTVMonitorQueueClient struct {
	grpc.ClientStream
}

func (x *jungleTVMonitorQueueClient) Recv() (*Queue, error) {
	m := new(Queue)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *jungleTVClient) RewardInfo(ctx context.Context, in *RewardInfoRequest, opts ...grpc.CallOption) (*RewardInfoResponse, error) {
	out := new(RewardInfoResponse)
	err := c.cc.Invoke(ctx, "/jungletv.JungleTV/RewardInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jungleTVClient) SubmitActivityChallenge(ctx context.Context, in *SubmitActivityChallengeRequest, opts ...grpc.CallOption) (*SubmitActivityChallengeResponse, error) {
	out := new(SubmitActivityChallengeResponse)
	err := c.cc.Invoke(ctx, "/jungletv.JungleTV/SubmitActivityChallenge", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jungleTVClient) ConsumeChat(ctx context.Context, in *ConsumeChatRequest, opts ...grpc.CallOption) (JungleTV_ConsumeChatClient, error) {
	stream, err := c.cc.NewStream(ctx, &_JungleTV_serviceDesc.Streams[4], "/jungletv.JungleTV/ConsumeChat", opts...)
	if err != nil {
		return nil, err
	}
	x := &jungleTVConsumeChatClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type JungleTV_ConsumeChatClient interface {
	Recv() (*ChatUpdate, error)
	grpc.ClientStream
}

type jungleTVConsumeChatClient struct {
	grpc.ClientStream
}

func (x *jungleTVConsumeChatClient) Recv() (*ChatUpdate, error) {
	m := new(ChatUpdate)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *jungleTVClient) SendChatMessage(ctx context.Context, in *SendChatMessageRequest, opts ...grpc.CallOption) (*SendChatMessageResponse, error) {
	out := new(SendChatMessageResponse)
	err := c.cc.Invoke(ctx, "/jungletv.JungleTV/SendChatMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jungleTVClient) SubmitProofOfWork(ctx context.Context, in *SubmitProofOfWorkRequest, opts ...grpc.CallOption) (*SubmitProofOfWorkResponse, error) {
	out := new(SubmitProofOfWorkResponse)
	err := c.cc.Invoke(ctx, "/jungletv.JungleTV/SubmitProofOfWork", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jungleTVClient) ForciblyEnqueueTicket(ctx context.Context, in *ForciblyEnqueueTicketRequest, opts ...grpc.CallOption) (*ForciblyEnqueueTicketResponse, error) {
	out := new(ForciblyEnqueueTicketResponse)
	err := c.cc.Invoke(ctx, "/jungletv.JungleTV/ForciblyEnqueueTicket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jungleTVClient) RemoveQueueEntry(ctx context.Context, in *RemoveQueueEntryRequest, opts ...grpc.CallOption) (*RemoveQueueEntryResponse, error) {
	out := new(RemoveQueueEntryResponse)
	err := c.cc.Invoke(ctx, "/jungletv.JungleTV/RemoveQueueEntry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jungleTVClient) RemoveChatMessage(ctx context.Context, in *RemoveChatMessageRequest, opts ...grpc.CallOption) (*RemoveChatMessageResponse, error) {
	out := new(RemoveChatMessageResponse)
	err := c.cc.Invoke(ctx, "/jungletv.JungleTV/RemoveChatMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jungleTVClient) SetChatSettings(ctx context.Context, in *SetChatSettingsRequest, opts ...grpc.CallOption) (*SetChatSettingsResponse, error) {
	out := new(SetChatSettingsResponse)
	err := c.cc.Invoke(ctx, "/jungletv.JungleTV/SetChatSettings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jungleTVClient) SetVideoEnqueuingEnabled(ctx context.Context, in *SetVideoEnqueuingEnabledRequest, opts ...grpc.CallOption) (*SetVideoEnqueuingEnabledResponse, error) {
	out := new(SetVideoEnqueuingEnabledResponse)
	err := c.cc.Invoke(ctx, "/jungletv.JungleTV/SetVideoEnqueuingEnabled", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jungleTVClient) BanUser(ctx context.Context, in *BanUserRequest, opts ...grpc.CallOption) (*BanUserResponse, error) {
	out := new(BanUserResponse)
	err := c.cc.Invoke(ctx, "/jungletv.JungleTV/BanUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jungleTVClient) RemoveBan(ctx context.Context, in *RemoveBanRequest, opts ...grpc.CallOption) (*RemoveBanResponse, error) {
	out := new(RemoveBanResponse)
	err := c.cc.Invoke(ctx, "/jungletv.JungleTV/RemoveBan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JungleTVServer is the server API for JungleTV service.
// All implementations must embed UnimplementedJungleTVServer
// for forward compatibility
type JungleTVServer interface {
	SignIn(*SignInRequest, JungleTV_SignInServer) error
	EnqueueMedia(context.Context, *EnqueueMediaRequest) (*EnqueueMediaResponse, error)
	MonitorTicket(*MonitorTicketRequest, JungleTV_MonitorTicketServer) error
	ConsumeMedia(*ConsumeMediaRequest, JungleTV_ConsumeMediaServer) error
	MonitorQueue(*MonitorQueueRequest, JungleTV_MonitorQueueServer) error
	RewardInfo(context.Context, *RewardInfoRequest) (*RewardInfoResponse, error)
	SubmitActivityChallenge(context.Context, *SubmitActivityChallengeRequest) (*SubmitActivityChallengeResponse, error)
	ConsumeChat(*ConsumeChatRequest, JungleTV_ConsumeChatServer) error
	SendChatMessage(context.Context, *SendChatMessageRequest) (*SendChatMessageResponse, error)
	SubmitProofOfWork(context.Context, *SubmitProofOfWorkRequest) (*SubmitProofOfWorkResponse, error)
	// moderation endpoints
	ForciblyEnqueueTicket(context.Context, *ForciblyEnqueueTicketRequest) (*ForciblyEnqueueTicketResponse, error)
	RemoveQueueEntry(context.Context, *RemoveQueueEntryRequest) (*RemoveQueueEntryResponse, error)
	RemoveChatMessage(context.Context, *RemoveChatMessageRequest) (*RemoveChatMessageResponse, error)
	SetChatSettings(context.Context, *SetChatSettingsRequest) (*SetChatSettingsResponse, error)
	SetVideoEnqueuingEnabled(context.Context, *SetVideoEnqueuingEnabledRequest) (*SetVideoEnqueuingEnabledResponse, error)
	BanUser(context.Context, *BanUserRequest) (*BanUserResponse, error)
	RemoveBan(context.Context, *RemoveBanRequest) (*RemoveBanResponse, error)
	mustEmbedUnimplementedJungleTVServer()
}

// UnimplementedJungleTVServer must be embedded to have forward compatible implementations.
type UnimplementedJungleTVServer struct {
}

func (UnimplementedJungleTVServer) SignIn(*SignInRequest, JungleTV_SignInServer) error {
	return status.Errorf(codes.Unimplemented, "method SignIn not implemented")
}
func (UnimplementedJungleTVServer) EnqueueMedia(context.Context, *EnqueueMediaRequest) (*EnqueueMediaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnqueueMedia not implemented")
}
func (UnimplementedJungleTVServer) MonitorTicket(*MonitorTicketRequest, JungleTV_MonitorTicketServer) error {
	return status.Errorf(codes.Unimplemented, "method MonitorTicket not implemented")
}
func (UnimplementedJungleTVServer) ConsumeMedia(*ConsumeMediaRequest, JungleTV_ConsumeMediaServer) error {
	return status.Errorf(codes.Unimplemented, "method ConsumeMedia not implemented")
}
func (UnimplementedJungleTVServer) MonitorQueue(*MonitorQueueRequest, JungleTV_MonitorQueueServer) error {
	return status.Errorf(codes.Unimplemented, "method MonitorQueue not implemented")
}
func (UnimplementedJungleTVServer) RewardInfo(context.Context, *RewardInfoRequest) (*RewardInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RewardInfo not implemented")
}
func (UnimplementedJungleTVServer) SubmitActivityChallenge(context.Context, *SubmitActivityChallengeRequest) (*SubmitActivityChallengeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitActivityChallenge not implemented")
}
func (UnimplementedJungleTVServer) ConsumeChat(*ConsumeChatRequest, JungleTV_ConsumeChatServer) error {
	return status.Errorf(codes.Unimplemented, "method ConsumeChat not implemented")
}
func (UnimplementedJungleTVServer) SendChatMessage(context.Context, *SendChatMessageRequest) (*SendChatMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendChatMessage not implemented")
}
func (UnimplementedJungleTVServer) SubmitProofOfWork(context.Context, *SubmitProofOfWorkRequest) (*SubmitProofOfWorkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitProofOfWork not implemented")
}
func (UnimplementedJungleTVServer) ForciblyEnqueueTicket(context.Context, *ForciblyEnqueueTicketRequest) (*ForciblyEnqueueTicketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ForciblyEnqueueTicket not implemented")
}
func (UnimplementedJungleTVServer) RemoveQueueEntry(context.Context, *RemoveQueueEntryRequest) (*RemoveQueueEntryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveQueueEntry not implemented")
}
func (UnimplementedJungleTVServer) RemoveChatMessage(context.Context, *RemoveChatMessageRequest) (*RemoveChatMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveChatMessage not implemented")
}
func (UnimplementedJungleTVServer) SetChatSettings(context.Context, *SetChatSettingsRequest) (*SetChatSettingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetChatSettings not implemented")
}
func (UnimplementedJungleTVServer) SetVideoEnqueuingEnabled(context.Context, *SetVideoEnqueuingEnabledRequest) (*SetVideoEnqueuingEnabledResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetVideoEnqueuingEnabled not implemented")
}
func (UnimplementedJungleTVServer) BanUser(context.Context, *BanUserRequest) (*BanUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BanUser not implemented")
}
func (UnimplementedJungleTVServer) RemoveBan(context.Context, *RemoveBanRequest) (*RemoveBanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveBan not implemented")
}
func (UnimplementedJungleTVServer) mustEmbedUnimplementedJungleTVServer() {}

// UnsafeJungleTVServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to JungleTVServer will
// result in compilation errors.
type UnsafeJungleTVServer interface {
	mustEmbedUnimplementedJungleTVServer()
}

func RegisterJungleTVServer(s grpc.ServiceRegistrar, srv JungleTVServer) {
	s.RegisterService(&_JungleTV_serviceDesc, srv)
}

func _JungleTV_SignIn_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SignInRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(JungleTVServer).SignIn(m, &jungleTVSignInServer{stream})
}

type JungleTV_SignInServer interface {
	Send(*SignInProgress) error
	grpc.ServerStream
}

type jungleTVSignInServer struct {
	grpc.ServerStream
}

func (x *jungleTVSignInServer) Send(m *SignInProgress) error {
	return x.ServerStream.SendMsg(m)
}

func _JungleTV_EnqueueMedia_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EnqueueMediaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JungleTVServer).EnqueueMedia(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jungletv.JungleTV/EnqueueMedia",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JungleTVServer).EnqueueMedia(ctx, req.(*EnqueueMediaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JungleTV_MonitorTicket_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(MonitorTicketRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(JungleTVServer).MonitorTicket(m, &jungleTVMonitorTicketServer{stream})
}

type JungleTV_MonitorTicketServer interface {
	Send(*EnqueueMediaTicket) error
	grpc.ServerStream
}

type jungleTVMonitorTicketServer struct {
	grpc.ServerStream
}

func (x *jungleTVMonitorTicketServer) Send(m *EnqueueMediaTicket) error {
	return x.ServerStream.SendMsg(m)
}

func _JungleTV_ConsumeMedia_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ConsumeMediaRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(JungleTVServer).ConsumeMedia(m, &jungleTVConsumeMediaServer{stream})
}

type JungleTV_ConsumeMediaServer interface {
	Send(*MediaConsumptionCheckpoint) error
	grpc.ServerStream
}

type jungleTVConsumeMediaServer struct {
	grpc.ServerStream
}

func (x *jungleTVConsumeMediaServer) Send(m *MediaConsumptionCheckpoint) error {
	return x.ServerStream.SendMsg(m)
}

func _JungleTV_MonitorQueue_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(MonitorQueueRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(JungleTVServer).MonitorQueue(m, &jungleTVMonitorQueueServer{stream})
}

type JungleTV_MonitorQueueServer interface {
	Send(*Queue) error
	grpc.ServerStream
}

type jungleTVMonitorQueueServer struct {
	grpc.ServerStream
}

func (x *jungleTVMonitorQueueServer) Send(m *Queue) error {
	return x.ServerStream.SendMsg(m)
}

func _JungleTV_RewardInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RewardInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JungleTVServer).RewardInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jungletv.JungleTV/RewardInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JungleTVServer).RewardInfo(ctx, req.(*RewardInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JungleTV_SubmitActivityChallenge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitActivityChallengeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JungleTVServer).SubmitActivityChallenge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jungletv.JungleTV/SubmitActivityChallenge",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JungleTVServer).SubmitActivityChallenge(ctx, req.(*SubmitActivityChallengeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JungleTV_ConsumeChat_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ConsumeChatRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(JungleTVServer).ConsumeChat(m, &jungleTVConsumeChatServer{stream})
}

type JungleTV_ConsumeChatServer interface {
	Send(*ChatUpdate) error
	grpc.ServerStream
}

type jungleTVConsumeChatServer struct {
	grpc.ServerStream
}

func (x *jungleTVConsumeChatServer) Send(m *ChatUpdate) error {
	return x.ServerStream.SendMsg(m)
}

func _JungleTV_SendChatMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendChatMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JungleTVServer).SendChatMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jungletv.JungleTV/SendChatMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JungleTVServer).SendChatMessage(ctx, req.(*SendChatMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JungleTV_SubmitProofOfWork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitProofOfWorkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JungleTVServer).SubmitProofOfWork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jungletv.JungleTV/SubmitProofOfWork",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JungleTVServer).SubmitProofOfWork(ctx, req.(*SubmitProofOfWorkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JungleTV_ForciblyEnqueueTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ForciblyEnqueueTicketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JungleTVServer).ForciblyEnqueueTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jungletv.JungleTV/ForciblyEnqueueTicket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JungleTVServer).ForciblyEnqueueTicket(ctx, req.(*ForciblyEnqueueTicketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JungleTV_RemoveQueueEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveQueueEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JungleTVServer).RemoveQueueEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jungletv.JungleTV/RemoveQueueEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JungleTVServer).RemoveQueueEntry(ctx, req.(*RemoveQueueEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JungleTV_RemoveChatMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveChatMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JungleTVServer).RemoveChatMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jungletv.JungleTV/RemoveChatMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JungleTVServer).RemoveChatMessage(ctx, req.(*RemoveChatMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JungleTV_SetChatSettings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetChatSettingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JungleTVServer).SetChatSettings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jungletv.JungleTV/SetChatSettings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JungleTVServer).SetChatSettings(ctx, req.(*SetChatSettingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JungleTV_SetVideoEnqueuingEnabled_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetVideoEnqueuingEnabledRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JungleTVServer).SetVideoEnqueuingEnabled(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jungletv.JungleTV/SetVideoEnqueuingEnabled",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JungleTVServer).SetVideoEnqueuingEnabled(ctx, req.(*SetVideoEnqueuingEnabledRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JungleTV_BanUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BanUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JungleTVServer).BanUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jungletv.JungleTV/BanUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JungleTVServer).BanUser(ctx, req.(*BanUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JungleTV_RemoveBan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveBanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JungleTVServer).RemoveBan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jungletv.JungleTV/RemoveBan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JungleTVServer).RemoveBan(ctx, req.(*RemoveBanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _JungleTV_serviceDesc = grpc.ServiceDesc{
	ServiceName: "jungletv.JungleTV",
	HandlerType: (*JungleTVServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EnqueueMedia",
			Handler:    _JungleTV_EnqueueMedia_Handler,
		},
		{
			MethodName: "RewardInfo",
			Handler:    _JungleTV_RewardInfo_Handler,
		},
		{
			MethodName: "SubmitActivityChallenge",
			Handler:    _JungleTV_SubmitActivityChallenge_Handler,
		},
		{
			MethodName: "SendChatMessage",
			Handler:    _JungleTV_SendChatMessage_Handler,
		},
		{
			MethodName: "SubmitProofOfWork",
			Handler:    _JungleTV_SubmitProofOfWork_Handler,
		},
		{
			MethodName: "ForciblyEnqueueTicket",
			Handler:    _JungleTV_ForciblyEnqueueTicket_Handler,
		},
		{
			MethodName: "RemoveQueueEntry",
			Handler:    _JungleTV_RemoveQueueEntry_Handler,
		},
		{
			MethodName: "RemoveChatMessage",
			Handler:    _JungleTV_RemoveChatMessage_Handler,
		},
		{
			MethodName: "SetChatSettings",
			Handler:    _JungleTV_SetChatSettings_Handler,
		},
		{
			MethodName: "SetVideoEnqueuingEnabled",
			Handler:    _JungleTV_SetVideoEnqueuingEnabled_Handler,
		},
		{
			MethodName: "BanUser",
			Handler:    _JungleTV_BanUser_Handler,
		},
		{
			MethodName: "RemoveBan",
			Handler:    _JungleTV_RemoveBan_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SignIn",
			Handler:       _JungleTV_SignIn_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "MonitorTicket",
			Handler:       _JungleTV_MonitorTicket_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ConsumeMedia",
			Handler:       _JungleTV_ConsumeMedia_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "MonitorQueue",
			Handler:       _JungleTV_MonitorQueue_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ConsumeChat",
			Handler:       _JungleTV_ConsumeChat_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "jungletv.proto",
}
