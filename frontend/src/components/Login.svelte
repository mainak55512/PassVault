<script>
  import { replace } from "svelte-spa-router";
  import {
    Login_cred,
    SetPin,
    IsNew,
    ChangePin,
    CheckSecurityQuestion,
    GenerateRSAKeys,
  } from "../../wailsjs/go/main/App.js";
  import logo from "../assets/images/padlock.png";
  import Swal from "sweetalert2";

  let in_value = "";

  function handleLogin() {
    Login_cred(in_value).then((result) => {
      if (in_value.length == 4 && result) {
        replace("/home");
      }
    });
  }

  async function storePin() {
    const steps = ["1", "2"];
    let new_pin = "";
    const Queue = Swal.mixin({
      progressSteps: steps,
      confirmButtonText: "Next &rarr;",
      // optional classes to avoid backdrop blinking between steps
      showClass: { backdrop: "swal2-noanimation" },
      hideClass: { backdrop: "swal2-noanimation" },
    });

    await Queue.fire({
      title: "Create PIN",
      html: `<input type="text" id="new_pin" class="swal2-input" maxlength="4" placeholder="New PIN">
  <input type="text" id="confirm_pin" class="swal2-input" maxlength="4" placeholder="Confirm PIN">`,
      // confirmButtonText: "Save &rarr;",
      focusConfirm: false,
      currentProgressStep: 0,
      background: "#303030",
      color: "white",
      preConfirm: () => {
        const newPin = Swal.getPopup().querySelector("#new_pin").value;
        const confirmPin = Swal.getPopup().querySelector("#confirm_pin").value;
        if (!newPin || !confirmPin) {
          Swal.showValidationMessage(`Please enter PIN`);
        }
        if (newPin !== confirmPin) {
          Swal.showValidationMessage(
            `New PIN and Confirm PIN must be the same`
          );
        }
        return { newPin: newPin };
      },
    }).then((result) => {
      new_pin = result.value.newPin;
    });
    await Queue.fire({
      title: "Security Question Setup",
      html: `<p>Please store the security questions as these might be required in password reset</p>
      <input type="text" id="ques1" class="swal2-input" placeholder="Favourite Singer">
  <input type="text" id="ques2" class="swal2-input" placeholder="Favourite Sport">`,
      confirmButtonText: "Save",
      focusConfirm: false,
      background: "#303030",
      color: "white",
      currentProgressStep: 1,
      preConfirm: () => {
        const ques1 = Swal.getPopup().querySelector("#ques1").value;
        const ques2 = Swal.getPopup().querySelector("#ques2").value;
        if (!ques1 || !ques2) {
          Swal.showValidationMessage(`Security questions cannot be empty!`);
        }
        return { ques1: ques1, ques2: ques2 };
      },
    }).then((result) => {
      SetPin(new_pin, result.value.ques1, result.value.ques2);
      GenerateRSAKeys();
      Swal.fire({
        text: "PIN created successfully!",
        icon: "success",
        background: "#303030",
        color: "white",
      });
    });
  }

  function changePin() {
    Swal.fire({
      title: "Change PIN",
      html: `<input type="text" id="old_pin" class="swal2-input" maxlength="4" placeholder="Old PIN">
  <input type="text" id="new_pin" class="swal2-input" maxlength="4" placeholder="New PIN">
  <input type="text" id="confirm_pin" class="swal2-input" maxlength="4" placeholder="Confirm new PIN">`,
      focusConfirm: false,
      currentProgressStep: 0,
      background: "#303030",
      color: "white",
      preConfirm: () => {
        const old_pin = Swal.getPopup().querySelector("#old_pin").value;
        const new_pin = Swal.getPopup().querySelector("#new_pin").value;
        const confirm_pin = Swal.getPopup().querySelector("#confirm_pin").value;
        let result = Login_cred(old_pin).then((res) => {
          return res;
        });

        if (!new_pin || !confirm_pin) {
          Swal.showValidationMessage(`Please enter PIN`);
        }
        if (new_pin !== confirm_pin) {
          Swal.showValidationMessage(
            `New PIN and Confirm PIN must be the same`
          );
        }
        return { new_pin: new_pin, valid_old_pin: result };
      },
    }).then((result) => {
      result.value.valid_old_pin.then((res) => {
        if (!res) {
          Swal.fire({
            icon: "error",
            title: "Invalid old pin",
            background: "#303030",
            color: "white",
          });
        } else {
          ChangePin(result.value.new_pin);
          Swal.fire({
            icon: "success",
            title: "Pin changed successfully!",
            background: "#303030",
            color: "white",
          });
        }
      });
    });
  }

  function createChangePin() {
    IsNew().then((isNew) => {
      if (isNew) {
        storePin();
      } else {
        changePin();
      }
    });
  }

  async function resetPin() {
    const steps = ["1", "2"];
    const Queue = Swal.mixin({
      progressSteps: steps,
      confirmButtonText: "Next &rarr;",
      // optional classes to avoid backdrop blinking between steps
      showClass: { backdrop: "swal2-noanimation" },
      hideClass: { backdrop: "swal2-noanimation" },
    });
    await Queue.fire({
      title: "Security Questions",
      html: `<input type="text" id="ques1" class="swal2-input" placeholder="Favourite Singer">
  <input type="text" id="ques2" class="swal2-input" placeholder="Favourite sport">`,
      // confirmButtonText: "Save &rarr;",
      focusConfirm: false,
      currentProgressStep: 0,
      background: "#303030",
      color: "white",
      preConfirm: () => {
        const ques1 = Swal.getPopup().querySelector("#ques1").value;
        const ques2 = Swal.getPopup().querySelector("#ques2").value;
        if (!ques1 || !ques2) {
          Swal.showValidationMessage(
            "Security questions are mandatory to fill"
          );
        }
        return { ques1: ques1, ques2: ques2 };
      },
    }).then((result) => {
      CheckSecurityQuestion(result.value.ques1, result.value.ques2).then(
        (res) => {
          if (!res) {
            Swal.fire({
              title: "Wrong answers!",
              icon: "error",
              background: "#303030",
              color: "white",
            });
          }
        }
      );
    });
    await Queue.fire({
      title: "Create PIN",
      html: `<input type="text" id="new_pin" class="swal2-input" maxlength="4" placeholder="New PIN">
  <input type="text" id="confirm_pin" class="swal2-input" maxlength="4" placeholder="Confirm PIN">`,
      // confirmButtonText: "Save &rarr;",
      focusConfirm: false,
      currentProgressStep: 1,
      background: "#303030",
      color: "white",
      preConfirm: () => {
        const newPin = Swal.getPopup().querySelector("#new_pin").value;
        const confirmPin = Swal.getPopup().querySelector("#confirm_pin").value;
        if (!newPin || !confirmPin) {
          Swal.showValidationMessage(`Please enter PIN`);
        }
        if (newPin !== confirmPin) {
          Swal.showValidationMessage(
            `New PIN and Confirm PIN must be the same`
          );
        }
        return { newPin: newPin };
      },
    }).then((result) => {
      ChangePin(result.value.newPin);
      Swal.fire({
        title: "PIN reset successful!",
        icon: "success",
        background: "#303030",
        color: "white",
      });
    });
  }
</script>

<div id="logo">
  <img alt="Password logo" src={logo} id="icon-style" />
  <div class="TitleText">PassVault</div>
</div>
<div class="pinbox">
  <p><b>Enter pin:</b></p>
  <input
    type="password"
    class="pin"
    maxlength="4"
    bind:value={in_value}
    on:input={() => handleLogin()}
  />
</div>
<br />
<div>
  <a
    href="javascript:void(0)"
    on:click={createChangePin}
    style="color:rgb(28, 245, 136)">Create/Change pin</a
  >
  │
  <a
    href="javascript:void(0)"
    on:click={resetPin}
    style="color:rgb(28, 245, 136)">Reset pin</a
  >
</div>

<style>
  .pin {
    border: none;
    background: transparent;
    border-bottom: 1px solid;
    border-color: aliceblue;
    text-align: center;
    font-size: 20px;
    color: antiquewhite;
    outline: none;
    letter-spacing: 30px;
    justify-content: center;
  }
  .pinbox {
    margin-top: 150px;
  }
  #logo {
    display: flex;
    flex-direction: column;
    width: 60px;
    height: 60px;
    margin: auto;
    margin-top: 80px;
    background-position: center;
    background-repeat: no-repeat;
    justify-content: center;
  }
  .TitleText {
    margin: auto;
    margin-top: 15px;
    font-weight: bold;
    margin-left: -10px;
  }
</style>
