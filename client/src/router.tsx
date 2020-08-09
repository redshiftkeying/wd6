import React, { Suspense } from "react";
import { HashRouter as Router, Route, Switch } from "react-router-dom";
import { useQuery, gql } from "@apollo/client";

export default function AdminRouter() {
  interface Routers {
    routers: Router[];
  }
  interface Router {
    path: string;
    exact: boolean;
    order: number;
    component: string;
  }
  const GET_DEFAULT_ROUTERS = gql`
    query routers {
      routers @cilent
    }
  `;
  const { data } = useQuery<Routers>(GET_DEFAULT_ROUTERS);

  return (
    <Router>
      {console.log("router hit:", data && data.routers)}
      <Suspense fallback={<div className="load">loading...</div>}>
        <Switch>
          {data &&
            data.routers.map((route, index) => (
              <Route
                key={index}
                path={route.path}
                exact={route.exact}
                component={React.lazy(() =>
                  import("./pages/" + route.component)
                )}
              />
            ))}
        </Switch>
      </Suspense>
    </Router>
  );
}
