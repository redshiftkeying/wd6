import { combineReducers } from 'redux'
import { TokenActionTypes, UserActionTypes, Router, RouterActionTypes } from './types'
import { TokenAction, UserAtion, RouterAction } from './actions'

export interface State { }

// token 
interface TokenState extends State {
    token: string
}

const initialTokenState: TokenState = {
    token: ""
}

function tokenReducer(
    state = initialTokenState,
    action: TokenAction
): TokenState {
    const newToken: TokenState = { token: "" }
    switch (action.type) {
        case TokenActionTypes.CREATE_TOKEN:
            newToken.token = action.token
            return newToken
        case TokenActionTypes.DELETE_TOKEN:
            return initialTokenState
        default: {
            return state
        }
    }
}

// user
interface UserState extends State {
    username: string
    showname: string
    password: string
    telphone: string
    avatar: string
    email: string
    isLogin: boolean
    msg: string
}

const initialUserState: UserState = {
    username: "",
    password: "",
    showname: "",
    telphone: "",
    avatar: "",
    email: "",
    isLogin: false,
    msg:"logging"
}

function userReducer(
    state = initialUserState,
    action: UserAtion
): UserState {
    let user: UserState = state
    switch (action.type) {
        case UserActionTypes.GET_USER:
            return { ...action.user, isLogin: true ,msg:action.msg}
        case UserActionTypes.USER_LOGOUT:
            return initialUserState
        case UserActionTypes.SET_USER_NAME:
            user.username = action.payload
            return user
        case UserActionTypes.SET_USER_PASSWORD:
            user.password = action.payload
            return user
        case UserActionTypes.LOGIN_FAIL:
            user.msg = action.msg
            return user
        default: {
            return state
        }
    }
}
// router tree
interface RouterState extends State {
    routers: Router[]
}
const initalRouterState: RouterState = {
    routers: [
        {
            order: 0,
            icon:"",
            path: "/",
            title: "",
            exact: true,
            component: "Login"
        },
        {
            order: 1,
            icon:"",
            path: "*",
            title: "404",
            exact: false,
            component: "404"
        }
    ]
}
function routerReducer(
    state = initalRouterState,
    action: RouterAction
): RouterState {
    switch (action.type) {
        case RouterActionTypes.GET_ROUTER:
            return state
        case RouterActionTypes.SET_ROUTER:
            return { routers: action.routers }
        default:
            return state
    }
}

export const rootReducer = combineReducers({
    token: tokenReducer,
    user: userReducer,
    router: routerReducer
})