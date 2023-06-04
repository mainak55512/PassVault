<script>
  import logo from "../assets/images/padlock.png";
  import { push } from "svelte-spa-router";

  function storepass() {
    push("/Store-password");
  }
  function fetchpass() {
    push("/Fetch-password");
  }
  function delpass() {
    push("/Delete-password");
  }
  function genpass() {
    push("/Generate-password");
  }
</script>

<div style="display: flex;" id="logo">
  <img alt="Password logo" src={logo} id="icon-style" />
  <div class="TitleText">PassVault</div>
</div>
<fieldset class="fieldset-style">
  <legend id="legend-style">Options</legend>
  <div class="btn-style">
    <button class="btn" on:click={fetchpass}>Fetch</button><br />
    <button class="btn btn-str" on:click={storepass}>Store</button><br />
    <button class="btn btn-del" on:click={delpass}>Delete</button><br />
    <button class="btn btn-gen" on:click={genpass}>Generate</button><br
    />
  </div>
</fieldset>

<style>
  #logo {
    display: block;
    width: 40px;
    height: 40px;
    margin: auto;
    margin-top: 80px;
    margin-bottom: 50px;
    background-position: center;
    background-repeat: no-repeat;
    background-size: 80% 50%;
    justify-content: center;
  }
  #icon-style {
    margin-right: 10px;
  }

  .TitleText {
    margin: 0.5rem auto;
    margin-top: 15px;
    font-weight: bold;
  }

  .btn {
    width: 200px;
    height: 50px;
    border-radius: 15px;
    border: 3px solid;
    margin: 0.6rem auto;
    padding: 0 8px;
    cursor: pointer;

    background: transparent;
    border-color: #1375f4;
    color: #1375f4;
    font-weight: bold;
    font-size: large;
  }

  .btn-str {
    border-color: #00c925;
    color: #00c925;
  }
  .btn-del {
    border-color: #fd1818;
    color: #fd1818;
  }
  .btn-gen {
    border-color: #fcfc46;
    color: #fcfc46;
  }

  .btn:hover {
    /* background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%); */
    background-color: #1375f4;
    color: white;
  }
  .btn.btn-str:hover {
    /* background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%); */
    background-color: #00c925;
    color: white;
  }
  .btn.btn-del:hover {
    /* background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%); */
    background-color: #fd1818;
    color: white;
  }
  .btn.btn-gen:hover {
    /* background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%); */
    background-color: #fcfc46;
    color: rgb(66, 65, 65);
  }

  .fieldset-style {
    width: 300px;
    margin: auto;
    padding: 5%;
    border-radius: 10px;
    text-align: left;
    border-color: gray;
    color: gray;
    border-width: 2px;
    border-style: solid;
  }
  .btn-style {
    text-align: center;
  }

  #legend-style {
    font-weight: bold;
    color: gray;
  }

  :global(body) {
    height: 100%;
    overflow-y: hidden;
  }
</style>
