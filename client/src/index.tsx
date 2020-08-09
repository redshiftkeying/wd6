import React from "react";
import ReactDOM from "react-dom";
import "./index.css";
import "./i18n";
import * as serviceWorker from "./serviceWorker";
import { Provider } from "react-redux";
import { ApolloProvider } from "@apollo/client";
import Router from "./router";
import store from "./store";
import client from "./graphql";

ReactDOM.render(
  <React.StrictMode>
    <Provider store={store}>
      <ApolloProvider client={client}>
        <Router />
      </ApolloProvider>
    </Provider>
  </React.StrictMode>,
  document.getElementById("root")
);

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA
serviceWorker.unregister();
