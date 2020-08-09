import { createStore, applyMiddleware, compose } from "redux";
import { rootReducer } from "./reducer";
import createSagaMiddleware from "redux-saga";
import iSagas from "./sagas";
const sagaMiddleware = createSagaMiddleware();

const composeEnhancer =
  (window as any).__REDUX_DEVTOOLS_EXTENSION_COMPOSE__ || compose;
const enhancer = composeEnhancer(applyMiddleware(sagaMiddleware));

const store = createStore(rootReducer, enhancer);
sagaMiddleware.run(iSagas);

export type RootState = ReturnType<typeof rootReducer>;
export default store;
