# License API

## Update Types

Run
```
just gen
```

## Add, remove, edit features
1. Add the features to `definitions/features.yaml`.
2. Run
```
just gen
```

## What happens when features are changed
Note: Any mention of "Stripe feature" is talking about a feature that has been created in Stripe. Any mention of "feature" without "Stripe"
immediately in front of it is referring to a feature defined in this repo's [features.yaml](./definitions/features.yaml).

Adding a feature will cause the CI to will create a corresponding Stripe feature. The name of the feature will be the Stripe feature's
`lookup-key` and the `displayName` will be the Stripe feature's `Name`. The Stripe feature's `Name` is the same as the feature's `displayName`
as a default.  If the feature's `preview` field is set to `true`, an additional feature will be created in Stripe. The additional "preview"
stripe Feature will be exactly like the normal Stripe feature, except its Stripe `lookup-key` will be appended with `-preview` and its Stripe
`Name` will be prefixed with `Preview: `.

A Stripe product will be created with a `Name` that matches the feature's `displayName`. Only the corresponding Stripe feature will be attached
to the Stripe product.

The Stripe feature will be attached to every existing Stripe product with the metadata key-value pair `attach_all_features=true`.

Changing the `displayName` name after a feature has a corresponding Stripe feature will have no effect on the Stripe feature. The feature's
`displayName` can be changed to reflect what a developer finds to be a good technical description. The Stripe feature's `Name` can be changed
to what the business side finds to be a good business description.

Changing the `name` of a feature is effectively the same as removing the feature, which has no effect in Stripe, and creating a new one with the
same fields, which would upload it to Stripe along with the additional Stripe behavior mentioned above.

Removing a feature will have no effect in Stripe. Stripe features exist indefinitely.

## Test Stripe feature upload CI locally
1. Create token in Stripe sandbox
2. Install `act`
```bash
curl --proto '=https' --tlsv1.2 -sSf https://raw.githubusercontent.com/nektos/act/master/install.sh | sudo bash
sudo mv bin/act /usr/local/bin
```
2. Run
```bahs
export STRIPE_API_KEY=<sandbox-token>
just upload-ci-local
```
