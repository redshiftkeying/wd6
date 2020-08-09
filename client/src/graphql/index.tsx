import {
  ApolloClient,
  InMemoryCache,
  createHttpLink,
  gql,
} from "@apollo/client";
import { setContext } from "@apollo/client/link/context";

const httpLink = createHttpLink({
  uri: "/graphql",
});

const authLink = setContext((_, { headers }) => {
  // get the authentication token from store if it exists
  const token = localStorage.getItem("token");
  // return the headers to the context so httpLink can read them
  return {
    headers: {
      ...headers,
      authorization: token ? `Bearer ${token}` : "",
    },
  };
});
const cache = new InMemoryCache();
// default loggin router by client
const DEFAULT_ROUTERS = gql`
  query routers {
    routers @cilent
  }
`;
// default router

cache.writeQuery({
  query: DEFAULT_ROUTERS,
  data: {
    routers: [
      {
        order: 0,
        icon: "",
        path: "/",
        title: "",
        exact: true,
        component: "Login",
      },
      {
        order: 1,
        icon: "",
        path: "*",
        title: "404",
        exact: false,
        component: "404",
      },
    ],
  },
});
const client = new ApolloClient({
  link: authLink.concat(httpLink),
  cache: cache,
});

export default client;
