<script>
	import { pop } from "svelte-spa-router";
	import { DelPass, FetchCredDetails } from "../../wailsjs/go/main/App.js";
	import Swal from "sweetalert2";
	let questions = [];
	let msg = "";

	FetchCredDetails().then((result) => {
		result.forEach((element) => {
			questions.push(element);
		});
		questions = questions;
	});

	let selected = "";

	function handleSubmit() {
		if (!selected) {
			msg = "**Please select the Service Name";
		} else {
			Swal.fire({
				title: "Are you sure?",
				text: "You won't be able to revert this!",
				icon: "warning",
				background: "#303030",
				color: "white",
				showCancelButton: true,
				confirmButtonColor: "#3085d6",
				cancelButtonColor: "#d33",
				confirmButtonText: "Yes, delete it!",
			}).then((result) => {
				if (result.isConfirmed) {
					DelPass(selected).then(() => {
						questions = questions.filter((el) => el != selected);
					});
					Swal.fire({
						title: "Deleted!",
						text: "Credentials has been deleted.",
						icon: "success",
						background: "#303030",
						color: "white",
					}).then(() => {
						selected = "";
					});
				}
			});
		}
	}
	function handleBack() {
		pop();
	}
</script>

<fieldset class="border-style">
	<legend class="legend-style"> Delete Details </legend>
	<div>
		<form on:submit|preventDefault={handleSubmit}>
			<select bind:value={selected} class="Selectbox-style">
				<option value="" disabled selected>
					Select The Service Name...</option
				>
				{#each questions as question}
					<option value={question}>
						{question}
					</option>
				{/each}
			</select>
			<br />
			<div class="btn-pos">
				<button type="submit" class="btn"> Delete </button>
				<button
					type="button"
					class="btn btn-extra"
					on:click={handleBack}
				>
					Back
				</button>
			</div>
		</form>
		<br />

		<p style="font-size: small; color: red">{msg}</p>
	</div>
</fieldset>

<style>
	*,
	*::before,
	*::after {
		box-sizing: border-box;
	}

	.btn-pos{
		display: flex;
	}

	.Selectbox-style {
		appearance: none;
		-webkit-appearance: none;
		-moz-appearance: none;
		background-color: rgb(71, 71, 91);
		height: 40px;
		padding-left: 10px;
		border: none;
		outline: none;
		margin: 0;
		margin-top: 70px;
		width: 300px;
		font-family: inherit;
		font-size: inherit;
		cursor: inherit;
		line-height: inherit;
		color: rgb(198, 196, 196);
		font-style: italic;
		font-family: "Trebuchet MS", "Lucida Sans Unicode", "Lucida Grande",
			"Lucida Sans", Arial, sans-serif;
		font-weight: normal;
		border-radius: 7px;
	}
	.btn {
		width: 100px;
		height: 40px;
		border-radius: 7px;
		border: none;
		margin: 2.5rem auto;
		margin-bottom: 1px;
		padding: 0 5px;
		cursor: pointer;
		background-color: transparent;
		color: whitesmoke;
		font-weight: bold;
		margin-left: 100px;
		margin-top:50px;
		border: 3px solid red;
		color: red;
		outline: none;
	}
	.btn-extra {
		background-color: transparent;
		margin: 0.5rem auto;
		margin-left: 40px;
		margin-top: 3.2rem;
		width: 50px;
		height: 40px;
		border: 3px solid rgb(27, 128, 228);
		color: rgb(27, 128, 228);
	}
	.border-style {
		width: 350px;
		height: auto;
		border-width: 2px;
		border-color: gray;
		border-radius: 10px;
		border-style: solid;
		margin: auto;
		margin-top: 140px;
		padding-left: 25px;
		display: block;
		text-align: left;
		padding-top: 10px;
	}
	.legend-style {
		color: gray;
		font-weight: bold;
	}
</style>
