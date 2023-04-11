package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["mtv/controllers:AuthController"] = append(beego.GlobalControllerRouter["mtv/controllers:AuthController"],
        beego.ControllerComments{
            Method: "CheckSign",
            Router: "/checksign",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mtv/controllers:GuardianController"] = append(beego.GlobalControllerRouter["mtv/controllers:GuardianController"],
        beego.ControllerComments{
            Method: "Add",
            Router: "/add",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mtv/controllers:GuardianController"] = append(beego.GlobalControllerRouter["mtv/controllers:GuardianController"],
        beego.ControllerComments{
            Method: "List",
            Router: "/list",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mtv/controllers:ImController"] = append(beego.GlobalControllerRouter["mtv/controllers:ImController"],
        beego.ControllerComments{
            Method: "CreateShareIm",
            Router: "/createshareim",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mtv/controllers:ImController"] = append(beego.GlobalControllerRouter["mtv/controllers:ImController"],
        beego.ControllerComments{
            Method: "ExchangeImPkey",
            Router: "/exchangeimpkey",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mtv/controllers:ImController"] = append(beego.GlobalControllerRouter["mtv/controllers:ImController"],
        beego.ControllerComments{
            Method: "Notify",
            Router: "/notify",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mtv/controllers:ImController"] = append(beego.GlobalControllerRouter["mtv/controllers:ImController"],
        beego.ControllerComments{
            Method: "Relays",
            Router: "/relays",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mtv/controllers:QuestionController"] = append(beego.GlobalControllerRouter["mtv/controllers:QuestionController"],
        beego.ControllerComments{
            Method: "Add",
            Router: "/add",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mtv/controllers:QuestionController"] = append(beego.GlobalControllerRouter["mtv/controllers:QuestionController"],
        beego.ControllerComments{
            Method: "List",
            Router: "/list",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mtv/controllers:QuestionController"] = append(beego.GlobalControllerRouter["mtv/controllers:QuestionController"],
        beego.ControllerComments{
            Method: "TmpList",
            Router: "/tmplist",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mtv/controllers:StorageController"] = append(beego.GlobalControllerRouter["mtv/controllers:StorageController"],
        beego.ControllerComments{
            Method: "Test",
            Router: "/test",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mtv/controllers:UserController"] = append(beego.GlobalControllerRouter["mtv/controllers:UserController"],
        beego.ControllerComments{
            Method: "BindMail",
            Router: "/bindmail",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mtv/controllers:UserController"] = append(beego.GlobalControllerRouter["mtv/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetImPubKeyList",
            Router: "/getimpubkeylist",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mtv/controllers:UserController"] = append(beego.GlobalControllerRouter["mtv/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetPassword",
            Router: "/getpassword",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mtv/controllers:UserController"] = append(beego.GlobalControllerRouter["mtv/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetSssData4Guardian",
            Router: "/getsssdata4guardian",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mtv/controllers:UserController"] = append(beego.GlobalControllerRouter["mtv/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetSssData4Question",
            Router: "/getsssdata4question",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mtv/controllers:UserController"] = append(beego.GlobalControllerRouter["mtv/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetUserInfo",
            Router: "/getuserinfo",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mtv/controllers:UserController"] = append(beego.GlobalControllerRouter["mtv/controllers:UserController"],
        beego.ControllerComments{
            Method: "ModifyUser",
            Router: "/modifyuser",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mtv/controllers:UserController"] = append(beego.GlobalControllerRouter["mtv/controllers:UserController"],
        beego.ControllerComments{
            Method: "SavePassword",
            Router: "/savepassword",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mtv/controllers:UserController"] = append(beego.GlobalControllerRouter["mtv/controllers:UserController"],
        beego.ControllerComments{
            Method: "SaveSssData4Guardian",
            Router: "/savesssdata4guardian",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mtv/controllers:UserController"] = append(beego.GlobalControllerRouter["mtv/controllers:UserController"],
        beego.ControllerComments{
            Method: "SaveSssData4Question",
            Router: "/savesssdata4question",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mtv/controllers:UserController"] = append(beego.GlobalControllerRouter["mtv/controllers:UserController"],
        beego.ControllerComments{
            Method: "SendMail4VerifyCode",
            Router: "/sendmail4verifycode",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mtv/controllers:UserController"] = append(beego.GlobalControllerRouter["mtv/controllers:UserController"],
        beego.ControllerComments{
            Method: "UpdateImgCid",
            Router: "/updateimgcid",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mtv/controllers:UserController"] = append(beego.GlobalControllerRouter["mtv/controllers:UserController"],
        beego.ControllerComments{
            Method: "UpdateImPkey",
            Router: "/updateimpkey",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mtv/controllers:UserController"] = append(beego.GlobalControllerRouter["mtv/controllers:UserController"],
        beego.ControllerComments{
            Method: "UpdateName",
            Router: "/updatename",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mtv/controllers:UserController"] = append(beego.GlobalControllerRouter["mtv/controllers:UserController"],
        beego.ControllerComments{
            Method: "UpdateSafeLevel",
            Router: "/updatesafelevel",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mtv/controllers:UserController"] = append(beego.GlobalControllerRouter["mtv/controllers:UserController"],
        beego.ControllerComments{
            Method: "VerifyMail",
            Router: "/verifymail",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
