<script>
    import { pop } from "svelte-spa-router";
    import { StorePass } from "../../wailsjs/go/main/App.js";
    import Swal from "sweetalert2";
    let service_name, user_id, password;
    let action_code = "";
    function savePass() {
        if (!service_name || !user_id || !password) {
            action_code = "**Cannot Submit... fields are empty";
        } else {
            let pwd = password;
            StorePass(service_name, user_id, pwd);
            Swal.fire({
                text: "Password saved successfully!",
                icon: "success",
                background: '#303030',
				color:'white',
            }).then(()=>{
                service_name = "";
                user_id = "";
                password = "";
            });
        }
    }
    function returnHome() {
        pop();
    }
</script>

<!-- <h3>Store Password</h3> -->

<fieldset class="fieldset-style">
    <legend class="legend-style"> Save Credentials </legend>
    <form class="content">
        <input
            type="text"
            placeholder="Service Name"
            bind:value={service_name}
            class="fields"
        />
        <input
            type="text"
            placeholder="User ID"
            bind:value={user_id}
            class="fields"
        />
        <input
            type="text"
            placeholder="Password"
            bind:value={password}
            class="fields"
        />
    </form>
    <div class="button-styles">
        <button class="btn" on:click={savePass}>Save</button>
        <button class="btn btn_home" on:click={returnHome}>back</button><br />
    </div>
    <p style="color: red;">{action_code}</p>
</fieldset>

<style>
    .fields {
        margin: 1rem auto;
        margin-left: -0.5rem;
        width: 250px;
        height: 40px;
        outline: none;
        /* border-top: none;
        border-left: none;
        border-right: none; */
        /* border-bottom: 0.3px solid white; */
        border: none;
        background: rgb(71, 71, 91);
        border-radius: 7px;
        padding-left: 10px;
        /* background-color: rgba(27, 38, 54, 1); */
        color: whitesmoke;
    }

    input::-webkit-input-placeholder {
        opacity: 0.7;
    }
    .button-styles {
        width: 300px;
        margin: auto;
    }

    .btn {
        width: 100px;
        height: 40px;
        border-radius: 7px;
        border: 3px solid rgb(28, 245, 136);
        margin: 0.6rem auto;
        margin-left: 18.5%;
        padding: 0 8px;
        cursor: pointer;
        background-color: transparent;
        color: rgb(28, 245, 136);
        font-weight: bold;
        margin-left: 75px;
    }
    .btn_home {
        width: 50px;
        height:40px;
        margin: 0.6rem 0 0.6rem 0;
        margin-right: 45px;
        float: inline-end;
        background-color: transparent;
        border-color: #1375f4;
        color: #1375f4;
    }
    .fieldset-style {
        margin: auto;
        margin-top: 130px;
        height: auto;
        width: 300px;
        border-radius: 10px;
        text-align: left;
        padding-left: 60px;
        padding-right: 0%;
        border-color: gray;
        border-style: solid;
    }
    .legend-style {
        color: gray;
        font-weight: bold;
        margin-left: -40px;
    }
</style>
