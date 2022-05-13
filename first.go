package main
import (
	"fmt"
	"rogchap.com/v8go"
)
func main() {
	  iso,_ := v8go.NewIsolate() 
	  ctx,_ := v8go.NewContext(iso) 
	  ctx.RunScript("class operation{ get(){ return 12 } }", "math.js") 
	  ctx.RunScript("var c = new operation()", "main.js") 
	  val1,_ := ctx.RunScript("c", "main1.js") 
	  ctx.RunScript("obj="+val1, "main2.js") 
	  ctx.RunScript("var r = obj.get();", "value.js") 
	  val, _ := ctx.RunScript("r", "value1.js")
	  fmt.Println(val) 
}


