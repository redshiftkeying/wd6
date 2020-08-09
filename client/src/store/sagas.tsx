import axios from 'axios'
import { createToken, getUser, userLoginFail } from './actions'
// import {GetRoutersAction} from './types'
import { put, takeEvery } from 'redux-saga/effects'
import { UserActionTypes } from './types'
import store from './index'

function* iSaga() {
    // yield all([
    //     tokenWatcher(),
    // ])
    yield takeEvery(UserActionTypes.USER_LOGIN, getToken)
}

function* getToken() {
    const State = store.getState()
    try {
        const res = yield axios.post("/login", {
            "username": State.user.username,
            "password": State.user.password
        })
        if (res.data.code === 0) {
            // login get token
            const tokenAction = createToken(res.data.data.token)
            const userAction = getUser(res.data.data.user, res.data.msg)
            // get user router
            // run
            yield put(tokenAction)
            yield put(userAction)
        } else {
            const failAction = userLoginFail(res.data.msg)
            yield put(failAction)
        }
    } catch (e) {
        const failAction = userLoginFail(e.message)
        yield put(failAction)
    }

}
// function* getRouters(){
//     const routers = yield axios.get('test/routers.json')
//     const action = GetRoutersAction(routers)
//     yield put(action)
// }
export default iSaga;