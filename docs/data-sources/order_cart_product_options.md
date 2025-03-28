---
subcategory : "Order"
---

# ovh_order_cart_product_options (Data Source)

Use this data source to retrieve information of order cart product options.

## Example Usage

```terraform
data "ovh_me" "my_account" {}

data "ovh_order_cart" "my_cart" {
  ovh_subsidiary = data.ovh_me.my_account.ovh_subsidiary
}

data "ovh_order_cart_product_options" "options" {
  cart_id   = data.ovh_order_cart.my_cart.id
  product   = "cloud"
  plan_code = "project"
}
```

## Argument Reference

* `cart_id` - (Required) Cart identifier
* `plan_code` - (Required) Product offer identifier
* `product` - (Required) Product
* `catalog_name` - Catalog name

## Attributes Reference

The following attributes are exported.

* `result` - products results
  * `plan_code` - Product offer identifier
  * `product_name` - Name of the product
  * `product_type` - Product type
  * `prices` - Prices of the product offer
    * `capacities` - Capacities of the pricing (type of pricing)
    * `description` - Description of the pricing
    * `duration` - Duration for ordering the product
    * `interval` - Interval of renewal
    * `maximum_quantity` - Maximum quantity that can be ordered
    * `maximum_repeat` - Maximum repeat for renewal
    * `minimum_quantity` - Minimum quantity that can be ordered
    * `minimum_repeat` - Minimum repeat for renewal
    * `price` - Price of the product (Price with its currency and textual representation)
      * `currency_code` - Currency code
      * `text` - Textual representation
      * `value` - The effective price
    * `price_in_ucents` - Price of the product in micro-centims
    * `pricing_mode` - Pricing model identifier
    * `pricing_type` - Pricing type
  * `exclusive` - Define if options of this family are exclusive with each other
  * `family` - Option family
  * `mandatory` - Define if an option of this family is mandatory
