// Code generated by csi-proxy-api-gen. DO NOT EDIT.

package v1beta2

import (
	v1beta2 "github.com/kubernetes-csi/csi-proxy/client/api/smb/v1beta2"
	internal "github.com/kubernetes-csi/csi-proxy/internal/server/smb/internal"
)

func autoConvert_v1beta2_NewSmbGlobalMappingRequest_To_internal_NewSmbGlobalMappingRequest(in *v1beta2.NewSmbGlobalMappingRequest, out *internal.NewSmbGlobalMappingRequest) error {
	out.RemotePath = in.RemotePath
	out.LocalPath = in.LocalPath
	out.Username = in.Username
	out.Password = in.Password
	return nil
}

// Convert_v1beta2_NewSmbGlobalMappingRequest_To_internal_NewSmbGlobalMappingRequest is an autogenerated conversion function.
func Convert_v1beta2_NewSmbGlobalMappingRequest_To_internal_NewSmbGlobalMappingRequest(in *v1beta2.NewSmbGlobalMappingRequest, out *internal.NewSmbGlobalMappingRequest) error {
	return autoConvert_v1beta2_NewSmbGlobalMappingRequest_To_internal_NewSmbGlobalMappingRequest(in, out)
}

func autoConvert_internal_NewSmbGlobalMappingRequest_To_v1beta2_NewSmbGlobalMappingRequest(in *internal.NewSmbGlobalMappingRequest, out *v1beta2.NewSmbGlobalMappingRequest) error {
	out.RemotePath = in.RemotePath
	out.LocalPath = in.LocalPath
	out.Username = in.Username
	out.Password = in.Password
	return nil
}

// Convert_internal_NewSmbGlobalMappingRequest_To_v1beta2_NewSmbGlobalMappingRequest is an autogenerated conversion function.
func Convert_internal_NewSmbGlobalMappingRequest_To_v1beta2_NewSmbGlobalMappingRequest(in *internal.NewSmbGlobalMappingRequest, out *v1beta2.NewSmbGlobalMappingRequest) error {
	return autoConvert_internal_NewSmbGlobalMappingRequest_To_v1beta2_NewSmbGlobalMappingRequest(in, out)
}

func autoConvert_v1beta2_NewSmbGlobalMappingResponse_To_internal_NewSmbGlobalMappingResponse(in *v1beta2.NewSmbGlobalMappingResponse, out *internal.NewSmbGlobalMappingResponse) error {
	out.Error = in.Error
	return nil
}

// Convert_v1beta2_NewSmbGlobalMappingResponse_To_internal_NewSmbGlobalMappingResponse is an autogenerated conversion function.
func Convert_v1beta2_NewSmbGlobalMappingResponse_To_internal_NewSmbGlobalMappingResponse(in *v1beta2.NewSmbGlobalMappingResponse, out *internal.NewSmbGlobalMappingResponse) error {
	return autoConvert_v1beta2_NewSmbGlobalMappingResponse_To_internal_NewSmbGlobalMappingResponse(in, out)
}

func autoConvert_internal_NewSmbGlobalMappingResponse_To_v1beta2_NewSmbGlobalMappingResponse(in *internal.NewSmbGlobalMappingResponse, out *v1beta2.NewSmbGlobalMappingResponse) error {
	out.Error = in.Error
	return nil
}

// Convert_internal_NewSmbGlobalMappingResponse_To_v1beta2_NewSmbGlobalMappingResponse is an autogenerated conversion function.
func Convert_internal_NewSmbGlobalMappingResponse_To_v1beta2_NewSmbGlobalMappingResponse(in *internal.NewSmbGlobalMappingResponse, out *v1beta2.NewSmbGlobalMappingResponse) error {
	return autoConvert_internal_NewSmbGlobalMappingResponse_To_v1beta2_NewSmbGlobalMappingResponse(in, out)
}

func autoConvert_v1beta2_RemoveSmbGlobalMappingRequest_To_internal_RemoveSmbGlobalMappingRequest(in *v1beta2.RemoveSmbGlobalMappingRequest, out *internal.RemoveSmbGlobalMappingRequest) error {
	out.RemotePath = in.RemotePath
	return nil
}

// Convert_v1beta2_RemoveSmbGlobalMappingRequest_To_internal_RemoveSmbGlobalMappingRequest is an autogenerated conversion function.
func Convert_v1beta2_RemoveSmbGlobalMappingRequest_To_internal_RemoveSmbGlobalMappingRequest(in *v1beta2.RemoveSmbGlobalMappingRequest, out *internal.RemoveSmbGlobalMappingRequest) error {
	return autoConvert_v1beta2_RemoveSmbGlobalMappingRequest_To_internal_RemoveSmbGlobalMappingRequest(in, out)
}

func autoConvert_internal_RemoveSmbGlobalMappingRequest_To_v1beta2_RemoveSmbGlobalMappingRequest(in *internal.RemoveSmbGlobalMappingRequest, out *v1beta2.RemoveSmbGlobalMappingRequest) error {
	out.RemotePath = in.RemotePath
	return nil
}

// Convert_internal_RemoveSmbGlobalMappingRequest_To_v1beta2_RemoveSmbGlobalMappingRequest is an autogenerated conversion function.
func Convert_internal_RemoveSmbGlobalMappingRequest_To_v1beta2_RemoveSmbGlobalMappingRequest(in *internal.RemoveSmbGlobalMappingRequest, out *v1beta2.RemoveSmbGlobalMappingRequest) error {
	return autoConvert_internal_RemoveSmbGlobalMappingRequest_To_v1beta2_RemoveSmbGlobalMappingRequest(in, out)
}

func autoConvert_v1beta2_RemoveSmbGlobalMappingResponse_To_internal_RemoveSmbGlobalMappingResponse(in *v1beta2.RemoveSmbGlobalMappingResponse, out *internal.RemoveSmbGlobalMappingResponse) error {
	out.Error = in.Error
	return nil
}

// Convert_v1beta2_RemoveSmbGlobalMappingResponse_To_internal_RemoveSmbGlobalMappingResponse is an autogenerated conversion function.
func Convert_v1beta2_RemoveSmbGlobalMappingResponse_To_internal_RemoveSmbGlobalMappingResponse(in *v1beta2.RemoveSmbGlobalMappingResponse, out *internal.RemoveSmbGlobalMappingResponse) error {
	return autoConvert_v1beta2_RemoveSmbGlobalMappingResponse_To_internal_RemoveSmbGlobalMappingResponse(in, out)
}

func autoConvert_internal_RemoveSmbGlobalMappingResponse_To_v1beta2_RemoveSmbGlobalMappingResponse(in *internal.RemoveSmbGlobalMappingResponse, out *v1beta2.RemoveSmbGlobalMappingResponse) error {
	out.Error = in.Error
	return nil
}

// Convert_internal_RemoveSmbGlobalMappingResponse_To_v1beta2_RemoveSmbGlobalMappingResponse is an autogenerated conversion function.
func Convert_internal_RemoveSmbGlobalMappingResponse_To_v1beta2_RemoveSmbGlobalMappingResponse(in *internal.RemoveSmbGlobalMappingResponse, out *v1beta2.RemoveSmbGlobalMappingResponse) error {
	return autoConvert_internal_RemoveSmbGlobalMappingResponse_To_v1beta2_RemoveSmbGlobalMappingResponse(in, out)
}