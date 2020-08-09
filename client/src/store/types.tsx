// Token Action types
export enum TokenActionTypes {
    CREATE_TOKEN = "getToken",
    DELETE_TOKEN = "delToken"
}

// User
export enum UserActionTypes {
    USER_LOGIN = "userLogin",
    LOGIN_FAIL="loginFail",
    GET_USER = "getUser",
    USER_LOGOUT = "userLogout",
    SET_USER_NAME = "setUserName",
    SET_USER_PASSWORD = "setUserPasword"
}
// User struct
export interface User {
    username: string
    showname: string
    password: string
    telphone: string
    avatar: string
    email: string
}
// Routers
export enum RouterActionTypes {
    GET_ROUTER = "getRouter",
    SET_ROUTER = "setRouter"
}

export interface Router {
    order: number
    icon: string
    path: string
    title: string
    exact: boolean
    component: string
}
