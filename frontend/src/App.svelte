<script>
  import Router from "svelte-spa-router";
  import Home from "./components/Home.svelte";
  import Store_password from "./components/Store-password.svelte";
  import Fetch_password from "./components/Fetch-password.svelte";
  import Password_details from "./components/Password-details.svelte";
  import Delete_password from "./components/Delete-password.svelte";
  import Generate_password from "./components/Generate-password.svelte";
  import Login from "./components/Login.svelte";

  const routes = {
    '/': Login,
    '/home':Home,
    '/Store-password':Store_password,
    '/Fetch-password':Fetch_password,
    '/Password-details/:service_name/:user_id/:pwd': Password_details,
    '/Delete-password':Delete_password,
    '/Generate-password':Generate_password
  }

</script>

<Router {routes}/>