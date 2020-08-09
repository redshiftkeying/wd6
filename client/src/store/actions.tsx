import { TokenActionTypes, RouterActionTypes, UserActionTypes, Router, User } from './types'

interface Action {
    type: string;
}

// token actions
export interface TokenAction extends Action {
    type: TokenActionTypes.DELETE_TOKEN | TokenActionTypes.CREATE_TOKEN
    token: string
}

export function createToken(newToken: string): TokenAction {
    return {
        type: TokenActionTypes.CREATE_TOKEN,
        token: newToken
    }
}

export function removeToken(): TokenAction {
    return {
        type: TokenActionTypes.DELETE_TOKEN,
        token: ""
    }
}

// user actions

interface UserLoginAction extends Action {
    type: UserActionTypes.USER_LOGIN
}

interface GetUserAction extends Action {
    type: UserActionTypes.GET_USER
    user: User
    msg: string
}

interface UserLogoutAction extends Action {
    type: UserActionTypes.USER_LOGOUT
}

interface UserLoginFailAction extends Action {
    type: UserActionTypes.LOGIN_FAIL
    msg: string
}
// User login 
export function userLogin(): UserLoginAction {
    return { type: UserActionTypes.USER_LOGIN }
}
export function removeUser(): UserLogoutAction {
    return { type: UserActionTypes.USER_LOGOUT }
}

export function userLoginFail(info: string): UserLoginFailAction {
    return {
        type: UserActionTypes.LOGIN_FAIL,
        msg: info
    }
}

export function getUser(newUser: User, msg: string): GetUserAction {
    return {
        type: UserActionTypes.GET_USER,
        user: newUser,
        msg: msg
    }
}
interface SetUserAction extends Action {
    type: UserActionTypes.SET_USER_NAME | UserActionTypes.SET_USER_PASSWORD
    payload: string
}
export function setUserName(name: string): SetUserAction {
    return {
        type: UserActionTypes.SET_USER_NAME,
        payload: name
    }
}
export function setUserPassword(pw: string): SetUserAction {
    return {
        type: UserActionTypes.SET_USER_PASSWORD,
        payload: pw
    }
}
export type UserAtion = UserLoginAction | UserLogoutAction | SetUserAction | GetUserAction | UserLoginFailAction

// routers actions
interface GetRouterAction extends Action {
    type: RouterActionTypes.GET_ROUTER
}
interface SetRouterAction extends Action {
    type: RouterActionTypes.SET_ROUTER
    routers: Router[]
}
export function getRouter(): GetRouterAction {
    return { type: RouterActionTypes.GET_ROUTER }
}
export function setRouter(rs: Router[]): SetRouterAction {
    return { type: RouterActionTypes.SET_ROUTER, routers: rs }
}
export type RouterAction = GetRouterAction | SetRouterAction