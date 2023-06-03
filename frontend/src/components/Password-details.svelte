<script>
    import { pop } from "svelte-spa-router";
    import { clickToCopy } from "../javascripts/copyToClipboard";
    export let params = {};
    function handleClick() {
        pop();
    }
</script>

<fieldset class="border-style">
    <legend style="color: rgb(184, 182, 182); font-weight:bold"
        >Credential Details</legend
    >
    <table class="space-under">
        <tbody>
            <tr>
                <th>Service Name:</th>
                <td class="sec-disp">{params.service_name}</td>
            </tr>
        </tbody>
        <tbody>
            <tr>
                <th>User ID:</th>
                <td class="sec-disp">{params.user_id}</td>
            </tr>
        </tbody>
        <tbody>
            <tr>
                <th>Password:</th>
                <td class="text-button"
                    ><p id="copy-text" class="sec-disp">{params.pwd}</p>
                    <button
                        use:clickToCopy={"#copy-text"}
                        class="cpy-btn"
                    /></td
                >
            </tr>
        </tbody>
    </table>
    <div class="btn_home">
        <button on:click={handleClick} class="back-btn">back</button>
    </div>
</fieldset>

<style>
    .sec-disp{
        background: rgb(71, 71, 91);
        border-radius: 7px;
        padding: 5px;
        font-weight: bold;
        word-break: break-all;
    }
    .border-style {
        width: 350px;
        height: auto;
        border-width: 2px;
        border-color: gray;
        border-radius: 7px;
        border-style: solid;
        margin: auto;
        margin-top: 110px;
        padding-left: 25px;
        display: block;
        text-align: left;
        padding-top: 10px;
    }
    .text-button {
        display: flex;
    }
    .btn_home {
        width: 350px;
        margin: auto;
        padding: 0;
        text-align: center;
    }
    .back-btn {
        /* float:; */
        margin-right: 0;
        margin-top: 5px;
        height: 40px;
        width: 50px;
        background-color: transparent;
        color: #1375f4;
        border-radius: 7px;
        border: 3px solid #1375f4;
        font-weight: bold;
    }
    .back-btn:hover{
        cursor:pointer;
    }
    .cpy-btn {
        height: 25px;
        width: 30px;
        margin: auto;
        margin-left: 5px;
        padding: 15px;
        border-radius: 7px;
        background: rgb(71, 71, 91);
        background-image: url(../assets/images/icons8-copy-24.png);
        background-repeat: no-repeat;
        background-position: center;
        background-size: 20px;
        border: none;
        outline: none;
        opacity: 0.8;
    }
    .cpy-btn:hover {
        cursor: pointer;
    }
    #copy-text {
        word-break: break-all;
    }
    .space-under {
        border-collapse: separate;
        border-spacing: 1.5em;
    }
</style>
