<script>
	import { push, pop } from "svelte-spa-router";
	import { FetchCredDetails, FetchPass } from "../../wailsjs/go/main/App.js";
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
			FetchPass(selected).then((result1) => {
				let pwd = encodeURIComponent(result1[2]);
				push(
					`/Password-details/${result1[0]}/${result1[1]}/${pwd}`
				);
			});
		}
	}
	function handleBack() {
		pop();
	}
</script>

<fieldset class="border-style">
	<legend class="legend-style"> Fetch Details </legend>
	<div>
			<form on:submit|preventDefault={handleSubmit}>
			<select bind:value={selected} class="Selectbox-style">
				<option value="" disabled> Select The Service Name...</option>
				{#each questions as question}
					<option value={question}>
						{question}
					</option>
				{/each}
			</select>
			<br />

			<div class="btn-pos">
			<button type="submit" class="btn"> Fetch </button>
			<button class="btn btn-extra" on:click={handleBack}> Back </button>
			</div>
		</form>

		<p style="font-size: small; color: red">{msg}</p>
	</div>
</fieldset>

<style>
	*,
	*::before,
	*::after {
		box-sizing: border-box;
	}

	.Selectbox-style {
		appearance: none;
		-webkit-appearance: none;
		-moz-appearance: none;
		background-color: rgb(71, 71, 91);
		border:none;
		outline: none;
		padding: 0 1em 0 0;
		margin: 0;
		margin-top: 70px;
		width: 300px;
		height: 40px;
		padding: 5px;
		padding-left: 10px;
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
		/* border: none; */
		border: 3px solid;
		border-color: rgb(28, 245, 136);
		margin: 2.5rem auto;
		margin-bottom: 20px;
		padding: 0 8px;
		cursor: pointer;
		background-color: transparent;
		color: rgb(28, 245, 136);
		font-weight: bold;
		margin-left: 100px;
		/* margin-right: auto; */
	}
	.btn-extra {
		background-color: transparent;
		border-color: #1375f4;
		color: #1375f4;
		/* margin: 0.5rem; */
		margin-left: 45px;
		/* margin-top:-00px; */
		width: 50px;
		height: 40px;
	}
	.btn-pos{
		display: flex;
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
