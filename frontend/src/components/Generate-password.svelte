<script>
    import { clickToCopy } from "../javascripts/copyToClipboard";
    import { pop } from "svelte-spa-router";
    import { GenPass } from "../../wailsjs/go/main/App.js";

    let pwd_gen = "";
    let pwd_len = 16;
    let upper_case= false,numeric = false,spl_char = false;

    function GeneratePassword(){
        GenPass(pwd_len, upper_case,numeric,spl_char).then((result)=>{
            pwd_gen = result;
        })
    }

    function handleBack(){
        pop();
    }

</script>

<fieldset class= "fieldset-style">
    <legend class="legend-style">Generate Password</legend>
<div class="pwd_section">
    <textarea id="text-style" readonly>{pwd_gen}</textarea>
    <button class="btn" use:clickToCopy={"#text-style"} />
</div>
<div class="slidecontainer">
    <!-- svelte-ignore a11y-label-has-associated-control -->
    <label id="label-style">Password length: </label>
    <input
        type="range"
        min="8"
        max="64"
        class="slider"
        id="myRange"
        bind:value={pwd_len}
    />
    <p id="sl_val">{pwd_len}</p>
</div>
<div class="toggle_case">
    <p style="font-weight: bold;font-size: 14px;">Include Uppercase letters:</p>
    <label class="switch">
        <input type="checkbox" bind:checked={upper_case}/>
        <span class="slider_1 round" />
    </label>
</div>
<div class="toggle_case toggle_case_number">
    <p style="font-weight: bold;font-size: 14px;">Include Numbers:</p>
    <label class="switch">
        <input type="checkbox" bind:checked={numeric}/>
        <span class="slider_1 round" />
    </label>
</div>
<div class="toggle_case toggle_case_spc_char">
    <p style="font-weight: bold;font-size: 14px;"> Include Special Characters:</p>
    <label class="switch">
        <input type="checkbox" bind:checked={spl_char}/>
        <span class="slider_1 round" />
    </label>
</div>

<div class="gen">
<button class="generate_btn" on:click={GeneratePassword}>Generate</button>
<button class="generate_btn back_btn" on:click={handleBack}>back</button>
</div>
</fieldset>

<style>
    .fieldset-style{
        margin:auto;
        margin-top: 100px;
        width: 350px;
        padding-bottom: 20px;
        border-radius: 10px;
        text-align: left;
        font-weight: bold;
        color: rgb(179, 163, 163);
        border: 2px solid gray;
    }
    .gen{
        padding-left: 65px;
    }
    .generate_btn{
        height: 40px;
        width: 80px;
        background: transparent;
        /* border: none; */
        border:3px solid;
        border-color: rgb(28, 245, 136);
        color: rgb(28, 245, 136);
        font-weight: bold;
        outline: none;
        border-radius: 10px;
        margin-top: 20px;
        margin-left: 75px;
    }
    .back_btn{
        width: 50px;
        background: transparent;
        border:3px solid;
        border-color: rgb(27, 128, 228);
        color: rgb(27, 128, 228);
        font-weight: bold;
        margin-left: 65px
    }
    .generate_btn,.back_btn:hover{
        cursor: pointer;
    }
    .toggle_case {
        display: flex;
        flex-direction: row;
        justify-content: center;
        margin-left: -78px;
    }
    .toggle_case_number{
        margin-left: -140px;
    }
    .toggle_case_spc_char{
        margin-left: -73.5px;
    }
    .pwd_section {
        display: flex;
        flex-direction: row;
        margin: auto;
        margin-top: 30px;
        text-align: left;
        color: white;
        background-color: rgb(71, 71, 91);
        width: 330px;
        height: 48px;
        overflow: hidden;
        border-radius: 7px;
    }
    #text-style {
        margin: auto;
        margin-left: 5px;
        font-size: 15px;
        width: 285px;
        height: 17px;
        word-break: break-all;
        background: transparent;
        resize: none;
        color: antiquewhite;
        border: none;
        outline: none;
    }
    .btn {
        margin: auto;
        text-align: end;
        background: rgb(31, 30, 40);
        background-image: url(../assets/images/icons8-copy-24.png);
        background-repeat: no-repeat;
        background-position: center;
        background-size: 20px;
        border: none;
        outline: none;
        opacity: 0.7;
        padding: 20px;
        border-radius: 7px;
        margin-right: 4px;
    }
    .slider {
        -webkit-appearance: none;
        width: 150px;
        height: 10px;
        border-radius: 5px;
        background: #d3d3d3;
        outline: none;
        opacity: 0.7;
        -webkit-transition: 0.2s;
        transition: opacity 0.2s;
        margin-top: 50px;
        margin-left: 7px;
    }

    .slider::-webkit-slider-thumb {
        -webkit-appearance: none;
        appearance: none;
        width: 20px;
        height: 20px;
        border-radius: 50%;
        background: rgb(28, 245, 136);
        border: 3px solid rgb(48, 139, 250);
        cursor: pointer;
    }

    .slider::-moz-range-thumb {
        width: 25px;
        height: 25px;
        border-radius: 50%;
        background: rgb(28, 245, 136);
        cursor: pointer;
    }
    #label-style {
        font-size: 14px;
        margin-top: 45px;
        font-weight: bold;
    }
    .slidecontainer {
        display: flex;
        flex-direction: row;
        justify-content: center;
    }
    #sl_val {
        margin-top: 42px;
        margin-left: 10px;
        height: 20px;
        widows: 30px;
        background: rgb(71, 71, 91);
        padding: 5px;
        border-radius: 7px;
    }
    .switch {
        position: relative;
        display: inline-block;
        width: 50px;
        height: 24px;
        margin-top: 12px;
        margin-left: 10px;
    }
    .switch input {
        opacity: 0;
        width: 0;
        height: 0;
    }
    .slider_1 {
        position: absolute;
        cursor: pointer;
        top: 0;
        left: 0;
        right: 0;
        bottom: 0;
        border: 2px solid;
        background: transparent;
        -webkit-transition: 0.4s;
        transition: 0.4s;
    }
    .slider_1:before {
        position: absolute;
        content: "";
        height: 12px;
        width: 12px;
        left: 2px;
        bottom: 1px;
        top: 2px;
        border: 2px solid;
        background: transparent;
        -webkit-transition: 0.4s;
        transition: 0.4s;
    }
    input:checked + .slider_1 {
        border-color: rgb(28, 245, 136);
    }
    input:checked + .slider_1:before {
        transform: translateX(26px);
    }
    .slider_1.round {
        border-radius: 34px;
    }
    .slider_1.round:before {
        border-radius: 50%;
    }
    /*

    */
</style>
