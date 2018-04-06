package eventbrite

import "golang.org/x/net/context"

// CategoriesResult is the response structure for the Categories
type CategoriesResult struct {
	Locale     string     `json:"locale"`
	Pagination Pagination `json:"pagination"`
	Categories []Category `json:"categories"`
}

// SubCategoriesResult is the response structure for the SubCategories
type SubCategoriesResult struct {
	Locale        string     `json:"locale"`
	Pagination    Pagination `json:"pagination"`
	Subcategories []Category `json:"subcategories"`
}

// Categories returns a list of category as categories, including subcategories nested
//
// https://www.eventbrite.com/developer/v3/endpoints/categories/#ebapi-get-categories
func (c *Client) Categories(ctx context.Context) (*CategoriesResult, error) {
	result := new(CategoriesResult)

	return result, c.getJSON(ctx, "/categories", nil, &result)
}

// Category gets a category by ID as category
//
// https://www.eventbrite.com/developer/v3/endpoints/categories/#ebapi-get-categories-id
func (c *Client) Category(ctx context.Context, id string) (*Category, error) {
	result := new(Category)

	return result, c.getJSON(ctx, "/categories/"+id, nil, &result)
}

// SubCategories gets a list of subcategory as subcategories
//
// https://www.eventbrite.com/developer/v3/endpoints/categories/#ebapi-get-subcategories
func (c *Client) SubCategories(ctx context.Context) (*SubCategoriesResult, error) {
	result := new(SubCategoriesResult)

	return result, c.getJSON(ctx, "/subcategories/", nil, &result)
}

// SubCategory gets a subcategory by ID as subcategory
//
// https://www.eventbrite.com/developer/v3/endpoints/categories/#ebapi-get-subcategories-id
func (c *Client) SubCategory(ctx context.Context, id string) (*SubCategory, error) {
	result := &SubCategory{}
	if err := c.getJSON(ctx, "/subcategories/"+id, nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}
