/*
 * Licensed to the OpenAirInterface (OAI) Software Alliance under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The OpenAirInterface Software Alliance licenses this file to You under
 * the Apache License, Version 2.0  (the "License"); you may not use this file
 * except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *-------------------------------------------------------------------------------
 * For more information about the OpenAirInterface (OAI) Software Alliance:
 *      contact@openairinterface.org
 */
#include <string.h>
#include <grpcpp/impl/codegen/status_code_enum.h>
#include <string>

#include "feg/protos/s6a_proxy.pb.h"

extern "C" {
#include "s6a_service_handler.h"
#include "log.h"
}
#include "S6aProxyImpl.h"

namespace grpc {
class Channel;
class ServerContext;
} // namespace grpc

using grpc::Channel;
using grpc::ServerContext;
using grpc::Status;
using grpc::StatusCode;

namespace magma {
using namespace feg;

S6aProxyImpl::S6aProxyImpl() {}

Status S6aProxyImpl::CancelLocation(
  ServerContext *context,
  const CancelLocationRequest *request,
  CancelLocationAnswer *answer)
{
  auto imsi = request->user_name();
  auto cancellation_type = request->cancellation_type();
  OAILOG_INFO(
    LOG_MME_APP,
    "Received CLR for %s of type %d\n ",
    imsi.c_str(),
    cancellation_type);
  if (cancellation_type == CancelLocationRequest::SUBSCRIPTION_WITHDRAWAL) {
    auto imsi_len = imsi.length();
    delete_subscriber_request(imsi.c_str(), imsi_len);
  }
  // return success regardless of cancellation type
  return Status::OK;
}

Status S6aProxyImpl::Reset(
  ServerContext *context,
  const ResetRequest *request,
  ResetAnswer *response)
{
  if (response != NULL) {
    response->set_error_code(ErrorCode::SUCCESS);
  }
  // Send message to MME_APP for furher processing
  handle_reset_request();
  // return success regardless of MME_APP processing status
  return Status::OK;
}

} // namespace magma
