package response

type Object struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// To transform the response (hide some paramters or add from anywhere else etc.), we may need this function.
// Right now i am just using it to convert model to json.

//func (o *Object) MapToResponse(model Model, ...) {
//	o.Key = model.key
//	o.Value = model.value
//		.
//		.
//		.
//	transformer logic here
//		.
//		.
//		.
//}
