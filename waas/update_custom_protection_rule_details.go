// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Web Application Acceleration and Security Services API
//
// OCI Web Application Acceleration and Security Services
//

package waas

import (
	"github.com/oracle/oci-go-sdk/v47/common"
)

// UpdateCustomProtectionRuleDetails Updates the configuration details of a custom protection rule. Custom protection rules can only be updated if they are not active in a WAAS policy.
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type UpdateCustomProtectionRuleDetails struct {

	// A user-friendly name for the custom protection rule.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// A description for the custom protection rule.
	Description *string `mandatory:"false" json:"description"`

	// The template text of the custom protection rule. All custom protection rules are expressed in ModSecurity Rule Language.
	// Additionally, each rule must include two placeholder variables that are updated by the WAF service upon publication of the rule.
	// `id: {{id_1}}` - This field is populated with a unique rule ID generated by the WAF service which identifies a `SecRule`. More than one `SecRule` can be defined in the `template` field of a CreateCustomSecurityRule call. The value of the first `SecRule` must be `id: {{id_1}}` and the `id` field of each subsequent `SecRule` should increase by one, as shown in the example.
	// `ctl:ruleEngine={{mode}}` - The action to be taken when the criteria of the `SecRule` are met, either `OFF`, `DETECT` or `BLOCK`. This field is automatically populated with the corresponding value of the `action` field of the `CustomProtectionRuleSetting` schema when the `WafConfig` is updated.
	// *Example:*
	//   ```
	//   SecRule REQUEST_COOKIES "regex matching SQL injection - part 1/2" \
	//           "phase:2,                                                 \
	//           msg:'Detects chained SQL injection attempts 1/2.',        \
	//           id: {{id_1}},                                             \
	//           ctl:ruleEngine={{mode}},                                  \
	//           deny"
	//   SecRule REQUEST_COOKIES "regex matching SQL injection - part 2/2" \
	//           "phase:2,                                                 \
	//           msg:'Detects chained SQL injection attempts 2/2.',        \
	//           id: {{id_2}},                                             \
	//           ctl:ruleEngine={{mode}},                                  \
	//           deny"
	//   ```
	//
	// The example contains two `SecRules` each having distinct regex expression to match the `Cookie` header value during the second input analysis phase.
	// For more information about custom protection rules, see Custom Protection Rules (https://docs.cloud.oracle.com/Content/WAF/tasks/customprotectionrules.htm).
	// For more information about ModSecurity syntax, see Making Rules: The Basic Syntax (https://www.modsecurity.org/CRS/Documentation/making.html).
	// For more information about ModSecurity's open source WAF rules, see Mod Security's OWASP Core Rule Set documentation (https://www.modsecurity.org/CRS/Documentation/index.html).
	Template *string `mandatory:"false" json:"template"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateCustomProtectionRuleDetails) String() string {
	return common.PointerString(m)
}
