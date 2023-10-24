// @generated by service-generator v0.0.1 with parameter "target=ts,import_extension=none,ts_nocheck=false"
// @generated from file github.com/rancher/opni/pkg/apis/management/v1/management.proto (package management, syntax proto3)
/* eslint-disable */

import { APIExtensionInfoList, CapabilityInstallerRequest, CapabilityInstallerResponse, CapabilityInstallRequest, CapabilityList, CapabilityStatusRequest, CapabilityUninstallCancelRequest, CapabilityUninstallRequest, CertsInfoResponse, CreateBootstrapTokenRequest, DashboardSettings, EditClusterRequest, GatewayConfig, ListClustersRequest, UpdateConfigRequest, WatchClustersRequest, WatchEvent } from "./management_pb";
import { AvailablePermissions, BackendRole, BackendRoleRequest, BootstrapToken, BootstrapTokenList, CapabilityType, CapabilityTypeList, Cluster, ClusterHealthStatus, ClusterList, HealthStatus, Reference, Role, RoleBinding, RoleBindingList, RoleList, TaskStatus } from "../../core/v1/core_pb";
import { axios } from "@pkg/opni/utils/axios";
import { Socket } from "@pkg/opni/utils/socket";
import { EVENT_CONNECT_ERROR, EVENT_CONNECTED, EVENT_CONNECTING, EVENT_DISCONNECT_ERROR, EVENT_MESSAGE } from "@shell/utils/socket";
import { Empty } from "@bufbuild/protobuf";
import { CancelUninstallRequest, InstallRequest, InstallResponse, NodeCapabilityStatus, StatusRequest, UninstallRequest, UninstallStatusRequest } from "../../capability/v1/capability_pb";


export async function CreateBootstrapToken(input: CreateBootstrapTokenRequest): Promise<BootstrapToken> {
  try {
    
    if (input) {
      console.info('Here is the input for a request to Management-CreateBootstrapToken:', input);
    }
  
    const response = (await axios.request({
    transformResponse: resp => BootstrapToken.fromBinary(new Uint8Array(resp)),
      method: 'post',
      responseType: 'arraybuffer',
      headers: {
        'Content-Type': 'application/octet-stream',
        'Accept': 'application/octet-stream',
      },
      url: `/opni-api/Management/tokens`,
    data: input?.toBinary() as ArrayBuffer
    })).data;

    console.info('Here is the response for a request to Management-CreateBootstrapToken:', response);
    return response
  } catch (ex: any) {
    if (ex?.response?.data) {
      const s = String.fromCharCode.apply(null, Array.from(new Uint8Array(ex?.response?.data)));
      console.error(s);
    }
    throw ex;
  }
}


export async function RevokeBootstrapToken(input: Reference): Promise<void> {
  try {
    
    if (input) {
      console.info('Here is the input for a request to Management-RevokeBootstrapToken:', input);
    }
  
    const response = (await axios.request({
      method: 'delete',
      responseType: 'arraybuffer',
      headers: {
        'Content-Type': 'application/octet-stream',
        'Accept': 'application/octet-stream',
      },
      url: `/opni-api/Management/tokens/${input.id}`,
    data: input?.toBinary() as ArrayBuffer
    })).data;

    console.info('Here is the response for a request to Management-RevokeBootstrapToken:', response);
    return response
  } catch (ex: any) {
    if (ex?.response?.data) {
      const s = String.fromCharCode.apply(null, Array.from(new Uint8Array(ex?.response?.data)));
      console.error(s);
    }
    throw ex;
  }
}


export async function ListBootstrapTokens(): Promise<BootstrapTokenList> {
  try {
    
    const response = (await axios.request({
    transformResponse: resp => BootstrapTokenList.fromBinary(new Uint8Array(resp)),
      method: 'get',
      responseType: 'arraybuffer',
      headers: {
        'Content-Type': 'application/octet-stream',
        'Accept': 'application/octet-stream',
      },
      url: `/opni-api/Management/tokens`
    })).data;

    console.info('Here is the response for a request to Management-ListBootstrapTokens:', response);
    return response
  } catch (ex: any) {
    if (ex?.response?.data) {
      const s = String.fromCharCode.apply(null, Array.from(new Uint8Array(ex?.response?.data)));
      console.error(s);
    }
    throw ex;
  }
}


export async function GetBootstrapToken(input: Reference): Promise<BootstrapToken> {
  try {
    
    if (input) {
      console.info('Here is the input for a request to Management-GetBootstrapToken:', input);
    }
  
    const response = (await axios.request({
    transformResponse: resp => BootstrapToken.fromBinary(new Uint8Array(resp)),
      method: 'get',
      responseType: 'arraybuffer',
      headers: {
        'Content-Type': 'application/octet-stream',
        'Accept': 'application/octet-stream',
      },
      url: `/opni-api/Management/tokens/${input.id}`,
    data: input?.toBinary() as ArrayBuffer
    })).data;

    console.info('Here is the response for a request to Management-GetBootstrapToken:', response);
    return response
  } catch (ex: any) {
    if (ex?.response?.data) {
      const s = String.fromCharCode.apply(null, Array.from(new Uint8Array(ex?.response?.data)));
      console.error(s);
    }
    throw ex;
  }
}


export async function ListClusters(input: ListClustersRequest): Promise<ClusterList> {
  try {
    
    if (input) {
      console.info('Here is the input for a request to Management-ListClusters:', input);
    }
  
    const response = (await axios.request({
    transformResponse: resp => ClusterList.fromBinary(new Uint8Array(resp)),
      method: 'get',
      responseType: 'arraybuffer',
      headers: {
        'Content-Type': 'application/octet-stream',
        'Accept': 'application/octet-stream',
      },
      url: `/opni-api/Management/clusters`,
    data: input?.toBinary() as ArrayBuffer
    })).data;

    console.info('Here is the response for a request to Management-ListClusters:', response);
    return response
  } catch (ex: any) {
    if (ex?.response?.data) {
      const s = String.fromCharCode.apply(null, Array.from(new Uint8Array(ex?.response?.data)));
      console.error(s);
    }
    throw ex;
  }
}


export function WatchClusters(input: WatchClustersRequest, callback: (data: WatchEvent) => void): () => Promise<any> {
  const socket = new Socket('/opni-api/Management/watch/clusters', true);
  Object.assign(socket, { frameTimeout: null })
  socket.addEventListener(EVENT_MESSAGE, (e: any) => {
    const event = e.detail;
    if (event.data) {
      callback(WatchEvent.fromBinary(new Uint8Array(event.data)));
    }
  });
  socket.addEventListener(EVENT_CONNECTING, () => {
    if (socket.socket) {
      socket.socket.binaryType = 'arraybuffer';
    }
  });
  socket.addEventListener(EVENT_CONNECTED, () => {
    socket.send(input.toBinary());
  });
  socket.addEventListener(EVENT_CONNECT_ERROR, (e) => {
    console.error(e);
  })
  socket.addEventListener(EVENT_DISCONNECT_ERROR, (e) => {
    console.error(e);
  })
  socket.connect();
  return () => {
    return socket.disconnect(null);
  };
}


export async function DeleteCluster(input: Reference): Promise<void> {
  try {
    
    if (input) {
      console.info('Here is the input for a request to Management-DeleteCluster:', input);
    }
  
    const response = (await axios.request({
      method: 'delete',
      responseType: 'arraybuffer',
      headers: {
        'Content-Type': 'application/octet-stream',
        'Accept': 'application/octet-stream',
      },
      url: `/opni-api/Management/clusters/${input.id}`,
    data: input?.toBinary() as ArrayBuffer
    })).data;

    console.info('Here is the response for a request to Management-DeleteCluster:', response);
    return response
  } catch (ex: any) {
    if (ex?.response?.data) {
      const s = String.fromCharCode.apply(null, Array.from(new Uint8Array(ex?.response?.data)));
      console.error(s);
    }
    throw ex;
  }
}


export async function CertsInfo(): Promise<CertsInfoResponse> {
  try {
    
    const response = (await axios.request({
    transformResponse: resp => CertsInfoResponse.fromBinary(new Uint8Array(resp)),
      method: 'get',
      responseType: 'arraybuffer',
      headers: {
        'Content-Type': 'application/octet-stream',
        'Accept': 'application/octet-stream',
      },
      url: `/opni-api/Management/certs`
    })).data;

    console.info('Here is the response for a request to Management-CertsInfo:', response);
    return response
  } catch (ex: any) {
    if (ex?.response?.data) {
      const s = String.fromCharCode.apply(null, Array.from(new Uint8Array(ex?.response?.data)));
      console.error(s);
    }
    throw ex;
  }
}


export async function GetCluster(input: Reference): Promise<Cluster> {
  try {
    
    if (input) {
      console.info('Here is the input for a request to Management-GetCluster:', input);
    }
  
    const response = (await axios.request({
    transformResponse: resp => Cluster.fromBinary(new Uint8Array(resp)),
      method: 'get',
      responseType: 'arraybuffer',
      headers: {
        'Content-Type': 'application/octet-stream',
        'Accept': 'application/octet-stream',
      },
      url: `/opni-api/Management/clusters/${input.id}`,
    data: input?.toBinary() as ArrayBuffer
    })).data;

    console.info('Here is the response for a request to Management-GetCluster:', response);
    return response
  } catch (ex: any) {
    if (ex?.response?.data) {
      const s = String.fromCharCode.apply(null, Array.from(new Uint8Array(ex?.response?.data)));
      console.error(s);
    }
    throw ex;
  }
}


export async function GetClusterHealthStatus(input: Reference): Promise<HealthStatus> {
  try {
    
    if (input) {
      console.info('Here is the input for a request to Management-GetClusterHealthStatus:', input);
    }
  
    const response = (await axios.request({
    transformResponse: resp => HealthStatus.fromBinary(new Uint8Array(resp)),
      method: 'get',
      responseType: 'arraybuffer',
      headers: {
        'Content-Type': 'application/octet-stream',
        'Accept': 'application/octet-stream',
      },
      url: `/opni-api/Management/clusters/${input.id}/health`,
    data: input?.toBinary() as ArrayBuffer
    })).data;

    console.info('Here is the response for a request to Management-GetClusterHealthStatus:', response);
    return response
  } catch (ex: any) {
    if (ex?.response?.data) {
      const s = String.fromCharCode.apply(null, Array.from(new Uint8Array(ex?.response?.data)));
      console.error(s);
    }
    throw ex;
  }
}


export function WatchClusterHealthStatus(input: Empty, callback: (data: ClusterHealthStatus) => void): () => Promise<any> {
  const socket = new Socket('/opni-api/Management/watch/health', true);
  Object.assign(socket, { frameTimeout: null })
  socket.addEventListener(EVENT_MESSAGE, (e: any) => {
    const event = e.detail;
    if (event.data) {
      callback(ClusterHealthStatus.fromBinary(new Uint8Array(event.data)));
    }
  });
  socket.addEventListener(EVENT_CONNECTING, () => {
    if (socket.socket) {
      socket.socket.binaryType = 'arraybuffer';
    }
  });
  socket.addEventListener(EVENT_CONNECTED, () => {
    socket.send(input.toBinary());
  });
  socket.addEventListener(EVENT_CONNECT_ERROR, (e) => {
    console.error(e);
  })
  socket.addEventListener(EVENT_DISCONNECT_ERROR, (e) => {
    console.error(e);
  })
  socket.connect();
  return () => {
    return socket.disconnect(null);
  };
}


export async function EditCluster(input: EditClusterRequest): Promise<Cluster> {
  try {
    
    if (input) {
      console.info('Here is the input for a request to Management-EditCluster:', input);
    }
  
    const response = (await axios.request({
    transformResponse: resp => Cluster.fromBinary(new Uint8Array(resp)),
      method: 'put',
      responseType: 'arraybuffer',
      headers: {
        'Content-Type': 'application/octet-stream',
        'Accept': 'application/octet-stream',
      },
      url: `/opni-api/Management/clusters/${input.cluster.id}`,
    data: input?.toBinary() as ArrayBuffer
    })).data;

    console.info('Here is the response for a request to Management-EditCluster:', response);
    return response
  } catch (ex: any) {
    if (ex?.response?.data) {
      const s = String.fromCharCode.apply(null, Array.from(new Uint8Array(ex?.response?.data)));
      console.error(s);
    }
    throw ex;
  }
}


export async function ListRBACBackends(): Promise<CapabilityTypeList> {
  try {
    return (await axios.request({
    transformResponse: resp => CapabilityTypeList.fromBinary(new Uint8Array(resp)),
      method: 'get',
      responseType: 'arraybuffer',
      headers: {
        'Content-Type': 'application/octet-stream',
        'Accept': 'application/octet-stream',
      },
      url: `/opni-api/Management/rbac/backend`
    })).data;
  } catch (ex) {
    if (ex?.response?.data) {
      const s = String.fromCharCode.apply(null, new Uint8Array(ex?.response?.data));
      console.error(s);
    }
    throw ex;
  }
}


export async function GetAvailableBackendPermissions(input: CapabilityType): Promise<AvailablePermissions> {
  try {
    return (await axios.request({
    transformResponse: resp => AvailablePermissions.fromBinary(new Uint8Array(resp)),
      method: 'get',
      responseType: 'arraybuffer',
      headers: {
        'Content-Type': 'application/octet-stream',
        'Accept': 'application/octet-stream',
      },
      url: `/opni-api/Management/rbac/backend/${input.name}/permissions`,
    data: input?.toBinary() as ArrayBuffer
    })).data;

    console.info('Here is the response for a request to Management-CreateRole:', response);
    return response
  } catch (ex: any) {
    if (ex?.response?.data) {
      const s = String.fromCharCode.apply(null, Array.from(new Uint8Array(ex?.response?.data)));
      console.error(s);
    }
    throw ex;
  }
}


export async function CreateBackendRole(input: BackendRole): Promise<void> {
  try {
    return (await axios.request({
      method: 'post',
      responseType: 'arraybuffer',
      headers: {
        'Content-Type': 'application/octet-stream',
        'Accept': 'application/octet-stream',
      },
      url: `/opni-api/Management/rbac/backend/${input.capability.name}/roles`,
    data: input?.toBinary() as ArrayBuffer
    })).data;

    console.info('Here is the response for a request to Management-UpdateRole:', response);
    return response
  } catch (ex: any) {
    if (ex?.response?.data) {
      const s = String.fromCharCode.apply(null, Array.from(new Uint8Array(ex?.response?.data)));
      console.error(s);
    }
    throw ex;
  }
}


export async function UpdateBackendRole(input: BackendRole): Promise<void> {
  try {
    return (await axios.request({
      method: 'put',
      responseType: 'arraybuffer',
      headers: {
        'Content-Type': 'application/octet-stream',
        'Accept': 'application/octet-stream',
      },
      url: `/opni-api/Management/rbac/backend/${input.capability.name}/roles`,
    data: input?.toBinary() as ArrayBuffer
    })).data;

    console.info('Here is the response for a request to Management-DeleteRole:', response);
    return response
  } catch (ex: any) {
    if (ex?.response?.data) {
      const s = String.fromCharCode.apply(null, Array.from(new Uint8Array(ex?.response?.data)));
      console.error(s);
    }
    throw ex;
  }
}


export async function DeleteBackendRole(input: BackendRoleRequest): Promise<void> {
  try {
    return (await axios.request({
      method: 'delete',
      responseType: 'arraybuffer',
      headers: {
        'Content-Type': 'application/octet-stream',
        'Accept': 'application/octet-stream',
      },
      url: `/opni-api/Management/rbac/backend/${input.capability.name}/roles/${input.roleRef.id}`,
    data: input?.toBinary() as ArrayBuffer
    })).data;

    console.info('Here is the response for a request to Management-GetRole:', response);
    return response
  } catch (ex: any) {
    if (ex?.response?.data) {
      const s = String.fromCharCode.apply(null, Array.from(new Uint8Array(ex?.response?.data)));
      console.error(s);
    }
    throw ex;
  }
}


export async function GetBackendRole(input: BackendRoleRequest): Promise<Role> {
  try {
    return (await axios.request({
    transformResponse: resp => Role.fromBinary(new Uint8Array(resp)),
      method: 'get',
      responseType: 'arraybuffer',
      headers: {
        'Content-Type': 'application/octet-stream',
        'Accept': 'application/octet-stream',
      },
      url: `/opni-api/Management/rbac/backend/${input.capability.name}/roles/${input.roleRef.id}`,
    data: input?.toBinary() as ArrayBuffer
    })).data;

    console.info('Here is the response for a request to Management-CreateRoleBinding:', response);
    return response
  } catch (ex: any) {
    if (ex?.response?.data) {
      const s = String.fromCharCode.apply(null, Array.from(new Uint8Array(ex?.response?.data)));
      console.error(s);
    }
    throw ex;
  }
}


export async function ListBackendRoles(input: CapabilityType): Promise<RoleList> {
  try {
    return (await axios.request({
    transformResponse: resp => RoleList.fromBinary(new Uint8Array(resp)),
      method: 'get',
      responseType: 'arraybuffer',
      headers: {
        'Content-Type': 'application/octet-stream',
        'Accept': 'application/octet-stream',
      },
      url: `/opni-api/Management/rbac/backend/${input.name}/roles`,
    data: input?.toBinary() as ArrayBuffer
    })).data;

    console.info('Here is the response for a request to Management-UpdateRoleBinding:', response);
    return response
  } catch (ex: any) {
    if (ex?.response?.data) {
      const s = String.fromCharCode.apply(null, Array.from(new Uint8Array(ex?.response?.data)));
      console.error(s);
    }
    throw ex;
  }
}


export async function CreateRoleBinding(input: RoleBinding): Promise<void> {
  try {
    return (await axios.request({
      method: 'post',
      responseType: 'arraybuffer',
      headers: {
        'Content-Type': 'application/octet-stream',
        'Accept': 'application/octet-stream',
      },
      url: `/opni-api/Management/rolebindings`,
    data: input?.toBinary() as ArrayBuffer
    })).data;

    console.info('Here is the response for a request to Management-DeleteRoleBinding:', response);
    return response
  } catch (ex: any) {
    if (ex?.response?.data) {
      const s = String.fromCharCode.apply(null, Array.from(new Uint8Array(ex?.response?.data)));
      console.error(s);
    }
    throw ex;
  }
}


export async function UpdateRoleBinding(input: RoleBinding): Promise<void> {
  try {
    return (await axios.request({
      method: 'put',
      responseType: 'arraybuffer',
      headers: {
        'Content-Type': 'application/octet-stream',
        'Accept': 'application/octet-stream',
      },
      url: `/opni-api/Management/rolebindings`,
    data: input?.toBinary() as ArrayBuffer
    })).data;

    console.info('Here is the response for a request to Management-GetRoleBinding:', response);
    return response
  } catch (ex: any) {
    if (ex?.response?.data) {
      const s = String.fromCharCode.apply(null, Array.from(new Uint8Array(ex?.response?.data)));
      console.error(s);
    }
    throw ex;
  }
}


export async function DeleteRoleBinding(input: Reference): Promise<void> {
  try {
    return (await axios.request({
      method: 'delete',
      responseType: 'arraybuffer',
      headers: {
        'Content-Type': 'application/octet-stream',
        'Accept': 'application/octet-stream',
      },
      url: `/opni-api/Management/rolebindings/${input.id}`,
    data: input?.toBinary() as ArrayBuffer
    })).data;

    console.info('Here is the response for a request to Management-ListRoles:', response);
    return response
  } catch (ex: any) {
    if (ex?.response?.data) {
      const s = String.fromCharCode.apply(null, Array.from(new Uint8Array(ex?.response?.data)));
      console.error(s);
    }
    throw ex;
  }
}


export async function GetRoleBinding(input: Reference): Promise<RoleBinding> {
  try {
    return (await axios.request({
    transformResponse: resp => RoleBinding.fromBinary(new Uint8Array(resp)),
      method: 'get',
      responseType: 'arraybuffer',
      headers: {
        'Content-Type': 'application/octet-stream',
        'Accept': 'application/octet-stream',
      },
      url: `/opni-api/Management/rolebindings/${input.id}`,
    data: input?.toBinary() as ArrayBuffer
    })).data;

    console.info('Here is the response for a request to Management-ListRoleBindings:', response);
    return response
  } catch (ex: any) {
    if (ex?.response?.data) {
      const s = String.fromCharCode.apply(null, Array.from(new Uint8Array(ex?.response?.data)));
      console.error(s);
    }
    throw ex;
  }
}


export async function ListRoleBindings(): Promise<RoleBindingList> {
  try {
    return (await axios.request({
    transformResponse: resp => RoleBindingList.fromBinary(new Uint8Array(resp)),
      method: 'get',
      responseType: 'arraybuffer',
      headers: {
        'Content-Type': 'application/octet-stream',
        'Accept': 'application/octet-stream',
      },
      url: `/opni-api/Management/rolebindings`
    })).data;

    console.info('Here is the response for a request to Management-SubjectAccess:', response);
    return response
  } catch (ex: any) {
    if (ex?.response?.data) {
      const s = String.fromCharCode.apply(null, Array.from(new Uint8Array(ex?.response?.data)));
      console.error(s);
    }
    throw ex;
  }
}


export async function APIExtensions(): Promise<APIExtensionInfoList> {
  try {
    
    const response = (await axios.request({
    transformResponse: resp => APIExtensionInfoList.fromBinary(new Uint8Array(resp)),
      method: 'get',
      responseType: 'arraybuffer',
      headers: {
        'Content-Type': 'application/octet-stream',
        'Accept': 'application/octet-stream',
      },
      url: `/opni-api/Management/apiextensions`
    })).data;

    console.info('Here is the response for a request to Management-APIExtensions:', response);
    return response
  } catch (ex: any) {
    if (ex?.response?.data) {
      const s = String.fromCharCode.apply(null, Array.from(new Uint8Array(ex?.response?.data)));
      console.error(s);
    }
    throw ex;
  }
}


export async function GetConfig(): Promise<GatewayConfig> {
  try {
    
    const response = (await axios.request({
    transformResponse: resp => GatewayConfig.fromBinary(new Uint8Array(resp)),
      method: 'get',
      responseType: 'arraybuffer',
      headers: {
        'Content-Type': 'application/octet-stream',
        'Accept': 'application/octet-stream',
      },
      url: `/opni-api/Management/config`
    })).data;

    console.info('Here is the response for a request to Management-GetConfig:', response);
    return response
  } catch (ex: any) {
    if (ex?.response?.data) {
      const s = String.fromCharCode.apply(null, Array.from(new Uint8Array(ex?.response?.data)));
      console.error(s);
    }
    throw ex;
  }
}


export async function UpdateConfig(input: UpdateConfigRequest): Promise<void> {
  try {
    
    if (input) {
      console.info('Here is the input for a request to Management-UpdateConfig:', input);
    }
  
    const response = (await axios.request({
      method: 'put',
      responseType: 'arraybuffer',
      headers: {
        'Content-Type': 'application/octet-stream',
        'Accept': 'application/octet-stream',
      },
      url: `/opni-api/Management/config`,
    data: input?.toBinary() as ArrayBuffer
    })).data;

    console.info('Here is the response for a request to Management-UpdateConfig:', response);
    return response
  } catch (ex: any) {
    if (ex?.response?.data) {
      const s = String.fromCharCode.apply(null, Array.from(new Uint8Array(ex?.response?.data)));
      console.error(s);
    }
    throw ex;
  }
}


export async function ListCapabilities(): Promise<CapabilityList> {
  try {
    
    const response = (await axios.request({
    transformResponse: resp => CapabilityList.fromBinary(new Uint8Array(resp)),
      method: 'get',
      responseType: 'arraybuffer',
      headers: {
        'Content-Type': 'application/octet-stream',
        'Accept': 'application/octet-stream',
      },
      url: `/opni-api/Management/capabilities`
    })).data;

    console.info('Here is the response for a request to Management-ListCapabilities:', response);
    return response
  } catch (ex: any) {
    if (ex?.response?.data) {
      const s = String.fromCharCode.apply(null, Array.from(new Uint8Array(ex?.response?.data)));
      console.error(s);
    }
    throw ex;
  }
}


export async function InstallCapability(input: InstallRequest): Promise<InstallResponse> {
  try {
    
    if (input) {
      console.info('Here is the input for a request to Management-InstallCapability:', input);
    }
  
    const response = (await axios.request({
    transformResponse: resp => InstallResponse.fromBinary(new Uint8Array(resp)),
      method: 'post',
      responseType: 'arraybuffer',
      headers: {
        'Content-Type': 'application/octet-stream',
        'Accept': 'application/octet-stream',
      },
      url: `/opni-api/Management/clusters/${input.agent.id}/capabilities/${input.capability.id}/install`,
    data: input?.toBinary() as ArrayBuffer
    })).data;

    console.info('Here is the response for a request to Management-InstallCapability:', response);
    return response
  } catch (ex: any) {
    if (ex?.response?.data) {
      const s = String.fromCharCode.apply(null, Array.from(new Uint8Array(ex?.response?.data)));
      console.error(s);
    }
    throw ex;
  }
}


export async function UninstallCapability(input: UninstallRequest): Promise<void> {
  try {
    
    if (input) {
      console.info('Here is the input for a request to Management-UninstallCapability:', input);
    }
  
    const response = (await axios.request({
      method: 'post',
      responseType: 'arraybuffer',
      headers: {
        'Content-Type': 'application/octet-stream',
        'Accept': 'application/octet-stream',
      },
      url: `/opni-api/Management/clusters/${input.agent.id}/capabilities/${input.capability.id}/uninstall`,
    data: input?.toBinary() as ArrayBuffer
    })).data;

    console.info('Here is the response for a request to Management-UninstallCapability:', response);
    return response
  } catch (ex: any) {
    if (ex?.response?.data) {
      const s = String.fromCharCode.apply(null, Array.from(new Uint8Array(ex?.response?.data)));
      console.error(s);
    }
    throw ex;
  }
}


export async function CapabilityStatus(input: StatusRequest): Promise<NodeCapabilityStatus> {
  try {
    
    if (input) {
      console.info('Here is the input for a request to Management-CapabilityStatus:', input);
    }
  
    const response = (await axios.request({
    transformResponse: resp => NodeCapabilityStatus.fromBinary(new Uint8Array(resp)),
      method: 'get',
      responseType: 'arraybuffer',
      headers: {
        'Content-Type': 'application/octet-stream',
        'Accept': 'application/octet-stream',
      },
      url: `/opni-api/Management/clusters/${input.agent.id}/capabilities/${input.capability.id}/status`,
    data: input?.toBinary() as ArrayBuffer
    })).data;

    console.info('Here is the response for a request to Management-CapabilityStatus:', response);
    return response
  } catch (ex: any) {
    if (ex?.response?.data) {
      const s = String.fromCharCode.apply(null, Array.from(new Uint8Array(ex?.response?.data)));
      console.error(s);
    }
    throw ex;
  }
}


export async function CapabilityUninstallStatus(input: UninstallStatusRequest): Promise<TaskStatus> {
  try {
    
    if (input) {
      console.info('Here is the input for a request to Management-CapabilityUninstallStatus:', input);
    }
  
    const response = (await axios.request({
    transformResponse: resp => TaskStatus.fromBinary(new Uint8Array(resp)),
      method: 'get',
      responseType: 'arraybuffer',
      headers: {
        'Content-Type': 'application/octet-stream',
        'Accept': 'application/octet-stream',
      },
      url: `/opni-api/Management/clusters/${input.agent.id}/capabilities/${input.capability.id}/uninstall/status`,
    data: input?.toBinary() as ArrayBuffer
    })).data;

    console.info('Here is the response for a request to Management-CapabilityUninstallStatus:', response);
    return response
  } catch (ex: any) {
    if (ex?.response?.data) {
      const s = String.fromCharCode.apply(null, Array.from(new Uint8Array(ex?.response?.data)));
      console.error(s);
    }
    throw ex;
  }
}


export async function CancelCapabilityUninstall(input: CancelUninstallRequest): Promise<void> {
  try {
    
    if (input) {
      console.info('Here is the input for a request to Management-CancelCapabilityUninstall:', input);
    }
  
    const response = (await axios.request({
      method: 'post',
      responseType: 'arraybuffer',
      headers: {
        'Content-Type': 'application/octet-stream',
        'Accept': 'application/octet-stream',
      },
      url: `/opni-api/Management/clusters/${input.agent.id}/capabilities/${input.capability.id}/uninstall/cancel`,
    data: input?.toBinary() as ArrayBuffer
    })).data;

    console.info('Here is the response for a request to Management-CancelCapabilityUninstall:', response);
    return response
  } catch (ex: any) {
    if (ex?.response?.data) {
      const s = String.fromCharCode.apply(null, Array.from(new Uint8Array(ex?.response?.data)));
      console.error(s);
    }
    throw ex;
  }
}


export async function GetDashboardSettings(): Promise<DashboardSettings> {
  try {
    
    const response = (await axios.request({
    transformResponse: resp => DashboardSettings.fromBinary(new Uint8Array(resp)),
      method: 'get',
      responseType: 'arraybuffer',
      headers: {
        'Content-Type': 'application/octet-stream',
        'Accept': 'application/octet-stream',
      },
      url: `/opni-api/Management/dashboard/settings`
    })).data;

    console.info('Here is the response for a request to Management-GetDashboardSettings:', response);
    return response
  } catch (ex: any) {
    if (ex?.response?.data) {
      const s = String.fromCharCode.apply(null, Array.from(new Uint8Array(ex?.response?.data)));
      console.error(s);
    }
    throw ex;
  }
}


export async function UpdateDashboardSettings(input: DashboardSettings): Promise<void> {
  try {
    
    if (input) {
      console.info('Here is the input for a request to Management-UpdateDashboardSettings:', input);
    }
  
    const response = (await axios.request({
      method: 'put',
      responseType: 'arraybuffer',
      headers: {
        'Content-Type': 'application/octet-stream',
        'Accept': 'application/octet-stream',
      },
      url: `/opni-api/Management/dashboard/settings`,
    data: input?.toBinary() as ArrayBuffer
    })).data;

    console.info('Here is the response for a request to Management-UpdateDashboardSettings:', response);
    return response
  } catch (ex: any) {
    if (ex?.response?.data) {
      const s = String.fromCharCode.apply(null, Array.from(new Uint8Array(ex?.response?.data)));
      console.error(s);
    }
    throw ex;
  }
}

