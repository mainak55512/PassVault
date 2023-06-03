import Swal from "sweetalert2";

export function clickToCopy(node, target) {
	async function copyText() {
		let text = target
			? document.querySelector(target).innerText
			: node.innerText;
		if (!text) {
			text = target
				? document.querySelector(target).value
				: node.value;
		}

		try {
			await navigator.clipboard.writeText(text);

			node.dispatchEvent(
				new CustomEvent('copysuccess', {
					bubbles: true
				})
			);
			Swal.fire({
				text: "Copied to clipboard!",
				icon: "success",
				showConfirmButton: false,
				background: '#303030',
				color: 'white',
				timer: 1200
			})
		} catch (error) {
			node.dispatchEvent(
				new CustomEvent('copyerror', {
					bubbles: true,
					detail: error
				})
			);
		}
	}

	node.addEventListener('click', copyText);



	return {
		destroy() {
			node.removeEventListener('click', copyText);
		}
	}
}