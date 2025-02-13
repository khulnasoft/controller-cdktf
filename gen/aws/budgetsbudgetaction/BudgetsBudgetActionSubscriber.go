package budgetsbudgetaction


type BudgetsBudgetActionSubscriber struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/budgets_budget_action#address BudgetsBudgetAction#address}.
	Address *string `field:"required" json:"address" yaml:"address"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/budgets_budget_action#subscription_type BudgetsBudgetAction#subscription_type}.
	SubscriptionType *string `field:"required" json:"subscriptionType" yaml:"subscriptionType"`
}

