/*
 * Copyright 2019 SAP SE or an SAP affiliate company. All rights reserved.
 * This file is licensed under the Apache Software License, v. 2 except as noted
 * otherwise in the LICENSE file
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 *
 */

package webhook

import (
	"net/http"

	"k8s.io/apimachinery/pkg/runtime"

	"github.com/gardener/controller-manager-library/pkg/controllermanager/cluster"
	"github.com/gardener/controller-manager-library/pkg/controllermanager/extension"
	areacfg "github.com/gardener/controller-manager-library/pkg/controllermanager/webhook/config"
	"github.com/gardener/controller-manager-library/pkg/logger"
)

////////////////////////////////////////////////////////////////////////////////
// Definition
////////////////////////////////////////////////////////////////////////////////
const CLUSTER_MAIN = "<MAIN>"

type WebhookKind string

const MUTATING = WebhookKind("mutating")
const VALIDATING = WebhookKind("validating")
const CONVERTING = WebhookKind("converting")

type Environment interface {
	extension.Environment
	GetConfig() *areacfg.Config

	RegisterWebhookByName(name string, target cluster.Interface) error
	RegisterWebhook(def Definition, target cluster.Interface) error
	RegisterWebhookGroup(name string, target cluster.Interface) error

	DeleteWebhookByName(name string, target cluster.Interface) error
	DeleteWebhook(def Definition, target cluster.Interface) error
}

type Interface interface {
	extension.ElementBase
	//admission.Interface

	GetEnvironment() Environment
	GetDefinition() Definition
	GetCluster() cluster.Interface
	GetScheme() *runtime.Scheme
}

type OptionDefinition extension.OptionDefinition

type Definition interface {
	GetName() string
	GetResources() []extension.ResourceKey
	GetScheme() *runtime.Scheme
	GetKind() WebhookKind
	GetHandler() WebhookHandler
	GetCluster() string
	ActivateExplicitly() bool

	ConfigOptions() map[string]OptionDefinition

	Definition() Definition
}

type WebhookHandler interface {
	GetKind() WebhookKind
	GetHTTPHandler(wh Interface) (http.Handler, error)

	String() string
}

type WebhookValidator interface {
	Validate(Interface) error
}

type HandlerFactory interface {
	CreateHandler() WebhookHandler
}

type RegistrationHandler interface {
	Kind() WebhookKind
	RegistrationResource() runtime.Object
	CreateDeclaration(def Definition, target cluster.Interface, client WebhookClientConfigSource) (WebhookDeclaration, error)
	Register(log logger.LogContext, labels map[string]string, cluster cluster.Interface, name string, declaration ...WebhookDeclaration) error
	Delete(name string, cluster cluster.Interface) error
}

type WebhookDeclaration interface {
	Kind() WebhookKind
}

type WebhookDeclarations []WebhookDeclaration

type ValidatorFunc func(Interface) error

func (this ValidatorFunc) Validate(wh Interface) error {
	return this(wh)
}
