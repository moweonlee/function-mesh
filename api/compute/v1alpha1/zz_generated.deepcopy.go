//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/api/autoscaling/v2beta2"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	autoscaling_k8s_iov1 "k8s.io/autoscaler/vertical-pod-autoscaler/pkg/apis/autoscaling.k8s.io/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthConfig) DeepCopyInto(out *AuthConfig) {
	*out = *in
	if in.OAuth2Config != nil {
		in, out := &in.OAuth2Config, &out.OAuth2Config
		*out = new(OAuth2Config)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthConfig.
func (in *AuthConfig) DeepCopy() *AuthConfig {
	if in == nil {
		return nil
	}
	out := new(AuthConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BatchSourceConfig) DeepCopyInto(out *BatchSourceConfig) {
	*out = *in
	if in.DiscoveryTriggererConfig != nil {
		in, out := &in.DiscoveryTriggererConfig, &out.DiscoveryTriggererConfig
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BatchSourceConfig.
func (in *BatchSourceConfig) DeepCopy() *BatchSourceConfig {
	if in == nil {
		return nil
	}
	out := new(BatchSourceConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Config.
func (in *Config) DeepCopy() *Config {
	if in == nil {
		return nil
	}
	out := new(Config)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsumerConfig) DeepCopyInto(out *ConsumerConfig) {
	*out = *in
	if in.SchemaProperties != nil {
		in, out := &in.SchemaProperties, &out.SchemaProperties
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ConsumerProperties != nil {
		in, out := &in.ConsumerProperties, &out.ConsumerProperties
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ReceiverQueueSize != nil {
		in, out := &in.ReceiverQueueSize, &out.ReceiverQueueSize
		*out = new(int32)
		**out = **in
	}
	if in.CryptoConfig != nil {
		in, out := &in.CryptoConfig, &out.CryptoConfig
		*out = new(CryptoConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsumerConfig.
func (in *ConsumerConfig) DeepCopy() *ConsumerConfig {
	if in == nil {
		return nil
	}
	out := new(ConsumerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CryptoConfig) DeepCopyInto(out *CryptoConfig) {
	*out = *in
	if in.CryptoKeyReaderConfig != nil {
		in, out := &in.CryptoKeyReaderConfig, &out.CryptoKeyReaderConfig
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.EncryptionKeys != nil {
		in, out := &in.EncryptionKeys, &out.EncryptionKeys
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.CryptoSecrets != nil {
		in, out := &in.CryptoSecrets, &out.CryptoSecrets
		*out = make([]CryptoSecret, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CryptoConfig.
func (in *CryptoConfig) DeepCopy() *CryptoConfig {
	if in == nil {
		return nil
	}
	out := new(CryptoConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CryptoSecret) DeepCopyInto(out *CryptoSecret) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CryptoSecret.
func (in *CryptoSecret) DeepCopy() *CryptoSecret {
	if in == nil {
		return nil
	}
	out := new(CryptoSecret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Function) DeepCopyInto(out *Function) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Function.
func (in *Function) DeepCopy() *Function {
	if in == nil {
		return nil
	}
	out := new(Function)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Function) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionList) DeepCopyInto(out *FunctionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Function, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionList.
func (in *FunctionList) DeepCopy() *FunctionList {
	if in == nil {
		return nil
	}
	out := new(FunctionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FunctionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionMesh) DeepCopyInto(out *FunctionMesh) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionMesh.
func (in *FunctionMesh) DeepCopy() *FunctionMesh {
	if in == nil {
		return nil
	}
	out := new(FunctionMesh)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FunctionMesh) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionMeshList) DeepCopyInto(out *FunctionMeshList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FunctionMesh, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionMeshList.
func (in *FunctionMeshList) DeepCopy() *FunctionMeshList {
	if in == nil {
		return nil
	}
	out := new(FunctionMeshList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FunctionMeshList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionMeshSpec) DeepCopyInto(out *FunctionMeshSpec) {
	*out = *in
	if in.Sources != nil {
		in, out := &in.Sources, &out.Sources
		*out = make([]SourceSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Sinks != nil {
		in, out := &in.Sinks, &out.Sinks
		*out = make([]SinkSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Functions != nil {
		in, out := &in.Functions, &out.Functions
		*out = make([]FunctionSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionMeshSpec.
func (in *FunctionMeshSpec) DeepCopy() *FunctionMeshSpec {
	if in == nil {
		return nil
	}
	out := new(FunctionMeshSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionMeshStatus) DeepCopyInto(out *FunctionMeshStatus) {
	*out = *in
	if in.SourceConditions != nil {
		in, out := &in.SourceConditions, &out.SourceConditions
		*out = make(map[string]ResourceCondition, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.SinkConditions != nil {
		in, out := &in.SinkConditions, &out.SinkConditions
		*out = make(map[string]ResourceCondition, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.FunctionConditions != nil {
		in, out := &in.FunctionConditions, &out.FunctionConditions
		*out = make(map[string]ResourceCondition, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionMeshStatus.
func (in *FunctionMeshStatus) DeepCopy() *FunctionMeshStatus {
	if in == nil {
		return nil
	}
	out := new(FunctionMeshStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionSpec) DeepCopyInto(out *FunctionSpec) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.MinReplicas != nil {
		in, out := &in.MinReplicas, &out.MinReplicas
		*out = new(int32)
		**out = **in
	}
	if in.MaxReplicas != nil {
		in, out := &in.MaxReplicas, &out.MaxReplicas
		*out = new(int32)
		**out = **in
	}
	in.Input.DeepCopyInto(&out.Input)
	in.Output.DeepCopyInto(&out.Output)
	if in.FuncConfig != nil {
		in, out := &in.FuncConfig, &out.FuncConfig
		*out = (*in).DeepCopy()
	}
	in.Resources.DeepCopyInto(&out.Resources)
	if in.SecretsMap != nil {
		in, out := &in.SecretsMap, &out.SecretsMap
		*out = make(map[string]SecretRef, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.VolumeMounts != nil {
		in, out := &in.VolumeMounts, &out.VolumeMounts
		*out = make([]v1.VolumeMount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AutoAck != nil {
		in, out := &in.AutoAck, &out.AutoAck
		*out = new(bool)
		**out = **in
	}
	if in.ForwardSourceMessageProperty != nil {
		in, out := &in.ForwardSourceMessageProperty, &out.ForwardSourceMessageProperty
		*out = new(bool)
		**out = **in
	}
	if in.MaxPendingAsyncRequests != nil {
		in, out := &in.MaxPendingAsyncRequests, &out.MaxPendingAsyncRequests
		*out = new(int32)
		**out = **in
	}
	in.Pod.DeepCopyInto(&out.Pod)
	if in.WindowConfig != nil {
		in, out := &in.WindowConfig, &out.WindowConfig
		*out = new(WindowConfig)
		(*in).DeepCopyInto(*out)
	}
	in.Messaging.DeepCopyInto(&out.Messaging)
	in.Runtime.DeepCopyInto(&out.Runtime)
	if in.StateConfig != nil {
		in, out := &in.StateConfig, &out.StateConfig
		*out = new(Stateful)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionSpec.
func (in *FunctionSpec) DeepCopy() *FunctionSpec {
	if in == nil {
		return nil
	}
	out := new(FunctionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionStatus) DeepCopyInto(out *FunctionStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(map[Component]ResourceCondition, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionStatus.
func (in *FunctionStatus) DeepCopy() *FunctionStatus {
	if in == nil {
		return nil
	}
	out := new(FunctionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GoRuntime) DeepCopyInto(out *GoRuntime) {
	*out = *in
	if in.Log != nil {
		in, out := &in.Log, &out.Log
		*out = new(RuntimeLogConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GoRuntime.
func (in *GoRuntime) DeepCopy() *GoRuntime {
	if in == nil {
		return nil
	}
	out := new(GoRuntime)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InputConf) DeepCopyInto(out *InputConf) {
	*out = *in
	if in.Topics != nil {
		in, out := &in.Topics, &out.Topics
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.CustomSerdeSources != nil {
		in, out := &in.CustomSerdeSources, &out.CustomSerdeSources
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.CustomSchemaSources != nil {
		in, out := &in.CustomSchemaSources, &out.CustomSchemaSources
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.SourceSpecs != nil {
		in, out := &in.SourceSpecs, &out.SourceSpecs
		*out = make(map[string]ConsumerConfig, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InputConf.
func (in *InputConf) DeepCopy() *InputConf {
	if in == nil {
		return nil
	}
	out := new(InputConf)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JavaRuntime) DeepCopyInto(out *JavaRuntime) {
	*out = *in
	if in.Log != nil {
		in, out := &in.Log, &out.Log
		*out = new(RuntimeLogConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.JavaOpts != nil {
		in, out := &in.JavaOpts, &out.JavaOpts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JavaRuntime.
func (in *JavaRuntime) DeepCopy() *JavaRuntime {
	if in == nil {
		return nil
	}
	out := new(JavaRuntime)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogConfig) DeepCopyInto(out *LogConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogConfig.
func (in *LogConfig) DeepCopy() *LogConfig {
	if in == nil {
		return nil
	}
	out := new(LogConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Messaging) DeepCopyInto(out *Messaging) {
	*out = *in
	if in.Pulsar != nil {
		in, out := &in.Pulsar, &out.Pulsar
		*out = new(PulsarMessaging)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Messaging.
func (in *Messaging) DeepCopy() *Messaging {
	if in == nil {
		return nil
	}
	out := new(Messaging)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OAuth2Config) DeepCopyInto(out *OAuth2Config) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OAuth2Config.
func (in *OAuth2Config) DeepCopy() *OAuth2Config {
	if in == nil {
		return nil
	}
	out := new(OAuth2Config)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OutputConf) DeepCopyInto(out *OutputConf) {
	*out = *in
	if in.ProducerConf != nil {
		in, out := &in.ProducerConf, &out.ProducerConf
		*out = new(ProducerConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.CustomSchemaSinks != nil {
		in, out := &in.CustomSchemaSinks, &out.CustomSchemaSinks
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OutputConf.
func (in *OutputConf) DeepCopy() *OutputConf {
	if in == nil {
		return nil
	}
	out := new(OutputConf)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodPolicy) DeepCopyInto(out *PodPolicy) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(v1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.SecurityContext != nil {
		in, out := &in.SecurityContext, &out.SecurityContext
		*out = new(v1.PodSecurityContext)
		(*in).DeepCopyInto(*out)
	}
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]v1.Volume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]v1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.InitContainers != nil {
		in, out := &in.InitContainers, &out.InitContainers
		*out = make([]v1.Container, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Sidecars != nil {
		in, out := &in.Sidecars, &out.Sidecars
		*out = make([]v1.Container, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.BuiltinAutoscaler != nil {
		in, out := &in.BuiltinAutoscaler, &out.BuiltinAutoscaler
		*out = make([]BuiltinHPARule, len(*in))
		copy(*out, *in)
	}
	if in.AutoScalingMetrics != nil {
		in, out := &in.AutoScalingMetrics, &out.AutoScalingMetrics
		*out = make([]v2beta2.MetricSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AutoScalingBehavior != nil {
		in, out := &in.AutoScalingBehavior, &out.AutoScalingBehavior
		*out = new(v2beta2.HorizontalPodAutoscalerBehavior)
		(*in).DeepCopyInto(*out)
	}
	if in.VPA != nil {
		in, out := &in.VPA, &out.VPA
		*out = new(VPASpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodPolicy.
func (in *PodPolicy) DeepCopy() *PodPolicy {
	if in == nil {
		return nil
	}
	out := new(PodPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProducerConfig) DeepCopyInto(out *ProducerConfig) {
	*out = *in
	if in.CryptoConfig != nil {
		in, out := &in.CryptoConfig, &out.CryptoConfig
		*out = new(CryptoConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProducerConfig.
func (in *ProducerConfig) DeepCopy() *ProducerConfig {
	if in == nil {
		return nil
	}
	out := new(ProducerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PulsarMessaging) DeepCopyInto(out *PulsarMessaging) {
	*out = *in
	if in.TLSConfig != nil {
		in, out := &in.TLSConfig, &out.TLSConfig
		*out = new(PulsarTLSConfig)
		**out = **in
	}
	if in.AuthConfig != nil {
		in, out := &in.AuthConfig, &out.AuthConfig
		*out = new(AuthConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PulsarMessaging.
func (in *PulsarMessaging) DeepCopy() *PulsarMessaging {
	if in == nil {
		return nil
	}
	out := new(PulsarMessaging)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PulsarStateStore) DeepCopyInto(out *PulsarStateStore) {
	*out = *in
	if in.JavaProvider != nil {
		in, out := &in.JavaProvider, &out.JavaProvider
		*out = new(PulsarStateStoreJavaProvider)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PulsarStateStore.
func (in *PulsarStateStore) DeepCopy() *PulsarStateStore {
	if in == nil {
		return nil
	}
	out := new(PulsarStateStore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PulsarStateStoreJavaProvider) DeepCopyInto(out *PulsarStateStoreJavaProvider) {
	*out = *in
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PulsarStateStoreJavaProvider.
func (in *PulsarStateStoreJavaProvider) DeepCopy() *PulsarStateStoreJavaProvider {
	if in == nil {
		return nil
	}
	out := new(PulsarStateStoreJavaProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PulsarTLSConfig) DeepCopyInto(out *PulsarTLSConfig) {
	*out = *in
	out.TLSConfig = in.TLSConfig
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PulsarTLSConfig.
func (in *PulsarTLSConfig) DeepCopy() *PulsarTLSConfig {
	if in == nil {
		return nil
	}
	out := new(PulsarTLSConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PythonRuntime) DeepCopyInto(out *PythonRuntime) {
	*out = *in
	if in.Log != nil {
		in, out := &in.Log, &out.Log
		*out = new(RuntimeLogConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PythonRuntime.
func (in *PythonRuntime) DeepCopy() *PythonRuntime {
	if in == nil {
		return nil
	}
	out := new(PythonRuntime)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceCondition) DeepCopyInto(out *ResourceCondition) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceCondition.
func (in *ResourceCondition) DeepCopy() *ResourceCondition {
	if in == nil {
		return nil
	}
	out := new(ResourceCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Runtime) DeepCopyInto(out *Runtime) {
	*out = *in
	if in.Java != nil {
		in, out := &in.Java, &out.Java
		*out = new(JavaRuntime)
		(*in).DeepCopyInto(*out)
	}
	if in.Python != nil {
		in, out := &in.Python, &out.Python
		*out = new(PythonRuntime)
		(*in).DeepCopyInto(*out)
	}
	if in.Golang != nil {
		in, out := &in.Golang, &out.Golang
		*out = new(GoRuntime)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Runtime.
func (in *Runtime) DeepCopy() *Runtime {
	if in == nil {
		return nil
	}
	out := new(Runtime)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuntimeLogConfig) DeepCopyInto(out *RuntimeLogConfig) {
	*out = *in
	if in.RotatePolicy != nil {
		in, out := &in.RotatePolicy, &out.RotatePolicy
		*out = new(TriggeringPolicy)
		**out = **in
	}
	if in.LogConfig != nil {
		in, out := &in.LogConfig, &out.LogConfig
		*out = new(LogConfig)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuntimeLogConfig.
func (in *RuntimeLogConfig) DeepCopy() *RuntimeLogConfig {
	if in == nil {
		return nil
	}
	out := new(RuntimeLogConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretRef) DeepCopyInto(out *SecretRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretRef.
func (in *SecretRef) DeepCopy() *SecretRef {
	if in == nil {
		return nil
	}
	out := new(SecretRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Sink) DeepCopyInto(out *Sink) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Sink.
func (in *Sink) DeepCopy() *Sink {
	if in == nil {
		return nil
	}
	out := new(Sink)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Sink) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SinkList) DeepCopyInto(out *SinkList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Sink, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SinkList.
func (in *SinkList) DeepCopy() *SinkList {
	if in == nil {
		return nil
	}
	out := new(SinkList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SinkList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SinkSpec) DeepCopyInto(out *SinkSpec) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.MinReplicas != nil {
		in, out := &in.MinReplicas, &out.MinReplicas
		*out = new(int32)
		**out = **in
	}
	if in.MaxReplicas != nil {
		in, out := &in.MaxReplicas, &out.MaxReplicas
		*out = new(int32)
		**out = **in
	}
	in.Input.DeepCopyInto(&out.Input)
	if in.SinkConfig != nil {
		in, out := &in.SinkConfig, &out.SinkConfig
		*out = (*in).DeepCopy()
	}
	in.Resources.DeepCopyInto(&out.Resources)
	if in.SecretsMap != nil {
		in, out := &in.SecretsMap, &out.SecretsMap
		*out = make(map[string]SecretRef, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.VolumeMounts != nil {
		in, out := &in.VolumeMounts, &out.VolumeMounts
		*out = make([]v1.VolumeMount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AutoAck != nil {
		in, out := &in.AutoAck, &out.AutoAck
		*out = new(bool)
		**out = **in
	}
	in.Pod.DeepCopyInto(&out.Pod)
	in.Messaging.DeepCopyInto(&out.Messaging)
	in.Runtime.DeepCopyInto(&out.Runtime)
	if in.StateConfig != nil {
		in, out := &in.StateConfig, &out.StateConfig
		*out = new(Stateful)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SinkSpec.
func (in *SinkSpec) DeepCopy() *SinkSpec {
	if in == nil {
		return nil
	}
	out := new(SinkSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SinkStatus) DeepCopyInto(out *SinkStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(map[Component]ResourceCondition, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SinkStatus.
func (in *SinkStatus) DeepCopy() *SinkStatus {
	if in == nil {
		return nil
	}
	out := new(SinkStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Source) DeepCopyInto(out *Source) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Source.
func (in *Source) DeepCopy() *Source {
	if in == nil {
		return nil
	}
	out := new(Source)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Source) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceList) DeepCopyInto(out *SourceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Source, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceList.
func (in *SourceList) DeepCopy() *SourceList {
	if in == nil {
		return nil
	}
	out := new(SourceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SourceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceSpec) DeepCopyInto(out *SourceSpec) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.MinReplicas != nil {
		in, out := &in.MinReplicas, &out.MinReplicas
		*out = new(int32)
		**out = **in
	}
	if in.MaxReplicas != nil {
		in, out := &in.MaxReplicas, &out.MaxReplicas
		*out = new(int32)
		**out = **in
	}
	in.Output.DeepCopyInto(&out.Output)
	if in.BatchSourceConfig != nil {
		in, out := &in.BatchSourceConfig, &out.BatchSourceConfig
		*out = new(BatchSourceConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.SourceConfig != nil {
		in, out := &in.SourceConfig, &out.SourceConfig
		*out = (*in).DeepCopy()
	}
	in.Resources.DeepCopyInto(&out.Resources)
	if in.SecretsMap != nil {
		in, out := &in.SecretsMap, &out.SecretsMap
		*out = make(map[string]SecretRef, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.VolumeMounts != nil {
		in, out := &in.VolumeMounts, &out.VolumeMounts
		*out = make([]v1.VolumeMount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ForwardSourceMessageProperty != nil {
		in, out := &in.ForwardSourceMessageProperty, &out.ForwardSourceMessageProperty
		*out = new(bool)
		**out = **in
	}
	in.Pod.DeepCopyInto(&out.Pod)
	in.Messaging.DeepCopyInto(&out.Messaging)
	in.Runtime.DeepCopyInto(&out.Runtime)
	if in.StateConfig != nil {
		in, out := &in.StateConfig, &out.StateConfig
		*out = new(Stateful)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceSpec.
func (in *SourceSpec) DeepCopy() *SourceSpec {
	if in == nil {
		return nil
	}
	out := new(SourceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceStatus) DeepCopyInto(out *SourceStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(map[Component]ResourceCondition, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceStatus.
func (in *SourceStatus) DeepCopy() *SourceStatus {
	if in == nil {
		return nil
	}
	out := new(SourceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Stateful) DeepCopyInto(out *Stateful) {
	*out = *in
	if in.Pulsar != nil {
		in, out := &in.Pulsar, &out.Pulsar
		*out = new(PulsarStateStore)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Stateful.
func (in *Stateful) DeepCopy() *Stateful {
	if in == nil {
		return nil
	}
	out := new(Stateful)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TLSConfig) DeepCopyInto(out *TLSConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TLSConfig.
func (in *TLSConfig) DeepCopy() *TLSConfig {
	if in == nil {
		return nil
	}
	out := new(TLSConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPASpec) DeepCopyInto(out *VPASpec) {
	*out = *in
	if in.UpdatePolicy != nil {
		in, out := &in.UpdatePolicy, &out.UpdatePolicy
		*out = new(autoscaling_k8s_iov1.PodUpdatePolicy)
		(*in).DeepCopyInto(*out)
	}
	if in.ResourcePolicy != nil {
		in, out := &in.ResourcePolicy, &out.ResourcePolicy
		*out = new(autoscaling_k8s_iov1.PodResourcePolicy)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPASpec.
func (in *VPASpec) DeepCopy() *VPASpec {
	if in == nil {
		return nil
	}
	out := new(VPASpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WindowConfig) DeepCopyInto(out *WindowConfig) {
	*out = *in
	if in.WindowLengthCount != nil {
		in, out := &in.WindowLengthCount, &out.WindowLengthCount
		*out = new(int32)
		**out = **in
	}
	if in.WindowLengthDurationMs != nil {
		in, out := &in.WindowLengthDurationMs, &out.WindowLengthDurationMs
		*out = new(int64)
		**out = **in
	}
	if in.SlidingIntervalCount != nil {
		in, out := &in.SlidingIntervalCount, &out.SlidingIntervalCount
		*out = new(int32)
		**out = **in
	}
	if in.SlidingIntervalDurationMs != nil {
		in, out := &in.SlidingIntervalDurationMs, &out.SlidingIntervalDurationMs
		*out = new(int64)
		**out = **in
	}
	if in.MaxLagMs != nil {
		in, out := &in.MaxLagMs, &out.MaxLagMs
		*out = new(int64)
		**out = **in
	}
	if in.WatermarkEmitIntervalMs != nil {
		in, out := &in.WatermarkEmitIntervalMs, &out.WatermarkEmitIntervalMs
		*out = new(int64)
		**out = **in
	}
	if in.TimestampExtractorClassName != nil {
		in, out := &in.TimestampExtractorClassName, &out.TimestampExtractorClassName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WindowConfig.
func (in *WindowConfig) DeepCopy() *WindowConfig {
	if in == nil {
		return nil
	}
	out := new(WindowConfig)
	in.DeepCopyInto(out)
	return out
}
