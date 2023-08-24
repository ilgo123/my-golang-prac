package productcontroller

import (
	"html/template"
	"net/http"
	"strconv"
	"time"
	homecontroller "web-native/controllers/homeController"
	"web-native/entities"
	"web-native/models/categorymodel"
	"web-native/models/productmodel"
)

func Index(w http.ResponseWriter, r *http.Request) {
	search := r.URL.Query().Get("search")
	if r.Method != http.MethodGet {
		response := entities.Response{
			Success: false,
			Message: "Method not allowed",
		}

		homecontroller.Response(w, http.StatusMethodNotAllowed, response)
		return
	}

	products := productmodel.GetAll(search)
	data := map[string]any{
		"products": products,
	}

	// temp, err := template.ParseFiles("views/product/index.html")
	// if err != nil {
	// 	panic(err)
	// }

	response := entities.Response{
		Success: true,
		Message: "Successfully get products data",
		Data:    data,
	}

	homecontroller.Response(w, http.StatusOK, response)
	// temp.Execute(w, data)
}
func Detail(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}

	product := productmodel.Detail(id)
	data := map[string]any{
		"product": product,
	}

	temp, err := template.ParseFiles("views/product/detail.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}
func Add(w http.ResponseWriter, r *http.Request) {
	// if r.Method == "GET" {
	// 	temp, err := template.ParseFiles("views/product/create.html")
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	categories := categorymodel.GetAll()
	// 	data := map[string]any{
	// 		"categories": categories,
	// 	}

	// 	temp.Execute(w, data)
	// }

	if r.Method == "POST" {
		var product entities.Product

		categoryId, err := strconv.Atoi(r.FormValue("category_id"))
		if err != nil {
			panic(err)
		}

		stock, err := strconv.Atoi(r.FormValue("stock"))
		if err != nil {
			panic(err)
		}

		product.Name = r.FormValue("name")
		product.Category.Id = uint(categoryId)
		product.Stock = int64(stock)
		product.Description = r.FormValue("description")
		product.CreatedAt = time.Now()
		product.UpdatedAt = time.Now()

		if ok := productmodel.Create(product); !ok {
			response := entities.Response{
				Success: false,
				Message: "Product failed to be created",
			}

			homecontroller.Response(w, http.StatusNotAcceptable, response)
			return
		}

		response := entities.Response{
			Success: true,
			Message: "Product successfully created",
		}

		homecontroller.Response(w, http.StatusCreated, response)
		// http.Redirect(w, r, "/products", http.StatusSeeOther)
	}
}
func Edit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/product/edit.html")
		if err != nil {
			panic(err)
		}

		idString := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		product := productmodel.Detail(id)

		search := r.URL.Query().Get("search")

		categories := categorymodel.GetAll(search)
		data := map[string]any{
			"categories": categories,
			"product":    product,
		}

		temp.Execute(w, data)
	}

	if r.Method == "POST" {
		var product entities.Product

		// id, err := strconv.Atoi(r.FormValue("id"))
		// if err != nil {
		// 	panic(err)
		// }

		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			panic(err)
		}

		categoryId, err := strconv.Atoi(r.FormValue("category_id"))
		if err != nil {
			panic(err)
		}

		stock, err := strconv.Atoi(r.FormValue("stock"))
		if err != nil {
			panic(err)
		}

		product.Name = r.FormValue("name")
		product.Category.Id = uint(categoryId)
		product.Stock = int64(stock)
		product.Description = r.FormValue("description")
		product.UpdatedAt = time.Now()

		if ok := productmodel.Update(id, product); !ok {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusTemporaryRedirect)
			return
		}

		http.Redirect(w, r, "/products", http.StatusSeeOther)
	}
}
func Delete(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		panic(err)
	}

	if err := productmodel.Delete(id); err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/products", http.StatusSeeOther)
}
