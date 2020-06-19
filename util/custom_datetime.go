package util

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strings"
	"time"
)

// DateTime overrides primitive.DateTime json marshalling and unmarshalling
type DateTime primitive.DateTime

func NewDateTimeFromTime(t time.Time) DateTime {
	return DateTime(primitive.NewDateTimeFromTime(t))
}

// UnmarshalJSON Parses the json string in the custom format
func (dt *DateTime) UnmarshalJSON(b []byte) (err error) {
	fmt.Println("UnmarshalJSON")
	s := strings.Trim(string(b), `"`)
	t, err := time.Parse(time.RFC3339, s)
	*dt = DateTime(primitive.NewDateTimeFromTime(t))
	return
}

// MarshalJSON writes a quoted string in the custom format
func (dt DateTime) MarshalJSON() ([]byte, error) {
	fmt.Println("MarshalJSON")
	return []byte(dt.String()), nil
}

// String returns the time in the custom format
func (dt *DateTime) String() string {
	t := primitive.DateTime(*dt).Time()
	return fmt.Sprintf("%q", t.Format(time.RFC3339))
}

// UnmarshalBSON Parses the json string in the custom format
func (dt DateTime) MarshalBSON() (interface{}, error) {
	fmt.Println("MarshalBSON")
	return primitive.DateTime(dt).Time(), nil
}

// func (d *DateTime) SetBSON(raw bson.Raw) error {
// 	return raw.Unmarshal(d)
// }

func (dt *DateTime) UnmarshalBSON(b []byte) error {
	fmt.Println("UnmarshalBSON")
	return bson.Unmarshal(b, (*primitive.DateTime)(dt).Time())
}

// // UnmarshalBSON Parses the json string in the custom format
// func (dt *DateTime) UnmarshalBSON(b []byte) (err error) {
// 	return bson.Unmarshal(b, dt)
// }

// // MarshalBSON writes a quoted string in the custom format
// func (dt DateTime) MarshalBSON() ([]byte, error) {
// 	return bson.Marshal(dt)
// }

// func (d DateTime) GetBSON() (interface{}, error) {
// 	return primitive.DateTime(d), nil
// }

// func (d *DateTime) SetBSON(raw bson.Raw) error {
// 	return primitive.DateTime(*d).SetBSON(raw)
// }

// func main() {
// 	t, err := time.Parse(time.RFC3339, "2020-06-11T17:04:06.000+05:30")
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return
// 	}

// 	fmt.Println(t.Format(time.RFC3339))
// 	dt := DateTime(primitive.NewDateTimeFromTime(t))
// 	fmt.Println(dt.String())

// 	b, e := dt.MarshalJSON()
// 	if e != nil {
// 		fmt.Println(e.Error())
// 		return
// 	}

// 	fmt.Println(string(b))
// 	pdt := &dt
// 	erro := pdt.UnmarshalJSON(b)
// 	if erro != nil {
// 		fmt.Println(erro)
// 	}
// 	fmt.Println(dt.String())

// }
