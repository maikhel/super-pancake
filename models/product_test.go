package models

import (
	"testing"

	. "github.com/franela/goblin"
	. "github.com/onsi/gomega"
	"gopkg.in/khaiql/dbcleaner.v2"
	"gopkg.in/khaiql/dbcleaner.v2/engine"
)

func Test(t *testing.T) {
	g := Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	Cleaner := dbcleaner.New()

	postgres := engine.NewPostgresEngine("postgres://" + DatabaseUser + "@" + DatabaseHost + "/" + DatabaseName + "?sslmode=disable")

	Cleaner.SetEngine(postgres)

	g.Describe("GetProducts", func() {
		g.BeforeEach(func() {
			Cleaner.Acquire("products")
		})

		g.AfterEach(func() {
			Cleaner.Clean("products")
		})

		g.It("Should return all products", func() {

			product := Product{
				Name:   "Banana",
				Amount: 5,
				Weight: 3,
			}

			err := DB.Create(&product).Error

			Expect(err).To(BeNil())

			otherProduct := Product{
				Name:   "Tomatoes",
				Amount: 5,
				Weight: 3,
			}

			err = DB.Create(&otherProduct).Error

			Expect(err).To(BeNil())

			processed, processedError := GetProducts()

			Expect(processedError).To(BeNil())

			Expect(len(*processed)).To(Equal(2))
			Expect((*processed)[0].Name).To(Equal("Banana"))
			Expect((*processed)[1].Name).To(Equal("Tomatoes"))
		})

		g.It("Should return empty array if no products", func() {

			processed, processedError := GetProducts()

			Expect(processedError).To(BeNil())

			Expect(len(*processed)).To(Equal(0))
		})
	})

	g.Describe("GetProduct", func() {
		g.BeforeEach(func() {
			Cleaner.Acquire("products")
		})

		g.AfterEach(func() {
			Cleaner.Clean("products")
		})

		g.It("Should find product by ID", func() {

			product := Product{
				Name:   "Banana",
				Amount: 5,
				Weight: 3,
			}

			err := DB.Create(&product).Error

			Expect(err).To(BeNil())

			processed, processedError := GetProduct(product.ID)

			Expect(processedError).To(BeNil())
			Expect(processed.Name).To(Equal("Banana"))
			Expect(product.ID).To(Equal(processed.ID))
			Expect(product.Name).To(Equal(processed.Name))
			Expect(product.Amount).To(Equal(processed.Amount))

		})

		g.It("Should return error if product not found", func() {

			_, processedError := GetProduct(7)

			Expect(processedError.Error()).To(Equal("record not found"))
		})
	})

	g.Describe("CreateProduct", func() {
		g.BeforeEach(func() {
			Cleaner.Acquire("products")
		})

		g.AfterEach(func() {
			Cleaner.Clean("products")
		})

		g.It("Should create a product", func() {
			input := Product{
				Name:   "Banana",
				Amount: 5,
				Weight: 3,
			}

			lastInput := Product{}

			processed, processedError := CreateProduct(&input)
			err := DB.Last(&lastInput).Error

			Expect(processedError).To(BeNil())
			Expect(processed.Name).To(Equal("Banana"))
			Expect(processed.Amount).To(Equal(5))

			Expect(err).To(BeNil())
			Expect(processed.ID).To(Equal(lastInput.ID))
		})
	})

	g.Describe("UpdateProduct", func() {
		g.BeforeEach(func() {
			Cleaner.Acquire("products")
		})

		g.AfterEach(func() {
			Cleaner.Clean("products")
		})

		g.It("Should update a product", func() {

			input := Product{
				Name:   "Banana",
				Amount: 5,
				Weight: 3,
			}

			DB.Create(&input)

			payload := map[string]interface{}{
				"Amount": 10,
			}

			processed, processedError := UpdateProduct(&input, payload)

			Expect(processedError).To(BeNil())
			Expect(processed.Amount).To(Equal(10))
		})
	})

	g.Describe("DeleteProduct", func() {
		g.BeforeEach(func() {
			Cleaner.Acquire("products")
		})

		g.AfterEach(func() {
			Cleaner.Clean("products")
		})

		g.It("Should delete a configuration", func() {
			product := Product{
				Name:   "Banana",
				Amount: 5,
				Weight: 3,
			}

			DB.Create(&product)

			_, processedError := DeleteProduct(&product)

			err := DB.Last(&product).Error

			Expect(processedError).To(BeNil())
			Expect(err.Error()).To(Equal("record not found"))
		})
	})
}
