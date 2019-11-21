package components

import (
	"html/template"
	"we/device_admin/models/usage"
)

const IconRaspberryPI = `<svg xmlns="http://www.w3.org/2000/svg" width="80" height="100" viewBox="0 0 48 48"><path fill="#37474F" d="M16.458,2.987c-0.212,0.007-0.441,0.086-0.701,0.29c-0.635-0.245-1.251-0.33-1.802,0.168c-0.852-0.109-1.128,0.118-1.337,0.383c-0.187-0.002-1.397-0.191-1.953,0.638C9.27,4.3,8.829,5.287,9.329,6.205c-0.285,0.441-0.58,0.877,0.086,1.719C9.179,8.394,9.325,8.9,9.88,9.516c-0.146,0.659,0.142,1.123,0.659,1.486c-0.097,0.9,0.826,1.424,1.102,1.61c0.106,0.526,0.326,1.021,1.38,1.295c0.174,0.784,0.808,0.917,1.421,1.083c-2.028,1.178-3.768,2.729-3.755,6.535l-0.297,0.529c-2.326,1.414-4.418,5.96-1.146,9.655c0.214,1.156,0.572,1.986,0.891,2.905c0.478,3.705,3.593,5.44,4.414,5.646c1.205,0.916,2.487,1.787,4.222,2.396c1.636,1.688,3.408,2.331,5.19,2.329c0.026,0,0.053,0.001,0.079,0c1.781,0.002,3.554-0.642,5.189-2.329c1.735-0.608,3.018-1.479,4.223-2.396c0.821-0.206,3.937-1.941,4.413-5.646c0.319-0.919,0.678-1.749,0.892-2.905c3.271-3.695,1.18-8.241-1.146-9.655l-0.297-0.53c0.012-3.805-1.729-5.356-3.756-6.534c0.613-0.166,1.247-0.3,1.421-1.084c1.055-0.272,1.275-0.769,1.381-1.295c0.276-0.186,1.198-0.709,1.103-1.611c0.517-0.361,0.805-0.826,0.657-1.484c0.557-0.615,0.702-1.124,0.466-1.592c0.667-0.842,0.371-1.277,0.087-1.719c0.499-0.918,0.059-1.905-1.337-1.739c-0.555-0.829-1.766-0.64-1.953-0.638c-0.209-0.265-0.486-0.492-1.337-0.383c-0.551-0.498-1.167-0.413-1.802-0.168c-0.756-0.596-1.256-0.119-1.826,0.062c-0.912-0.298-1.122,0.111-1.57,0.277c-0.997-0.211-1.299,0.247-1.777,0.731l-0.556-0.011c-1.503,0.886-2.249,2.69-2.514,3.616c-0.264-0.928-1.009-2.731-2.512-3.616l-0.556,0.011c-0.479-0.484-0.781-0.942-1.778-0.731c-0.448-0.166-0.657-0.575-1.571-0.277C17.208,3.22,16.863,2.975,16.458,2.987L16.458,2.987z"/><path fill="#64DD17" d="M13.466 6.885c3.987 2.055 6.305 3.718 7.575 5.134-.65 2.607-4.042 2.726-5.283 2.653.254-.119.467-.26.541-.479-.311-.221-1.415-.023-2.186-.456.296-.062.435-.12.573-.339-.727-.232-1.511-.433-1.973-.817.249.003.481.055.806-.17-.652-.351-1.348-.629-1.888-1.166.337-.009.701-.004.806-.129-.596-.37-1.1-.78-1.518-1.23.472.058.671.009.786-.075-.452-.461-1.023-.85-1.294-1.421.35.121.671.168.902-.011-.154-.345-.81-.55-1.189-1.357.369.036.761.081.839 0-.172-.697-.465-1.089-.753-1.496.79-.01 1.985.004 1.931-.063l-.488-.499c.771-.207 1.561.034 2.133.213.257-.203-.005-.459-.318-.721.655.087 1.247.238 1.782.445.286-.258-.186-.516-.413-.773 1.012.191 1.44.46 1.866.73.31-.295.018-.548-.19-.807.764.283 1.156.648 1.57 1.009.141-.19.357-.328.096-.784.542.312.95.68 1.252 1.092.335-.214.2-.506.201-.775.563.459.921.946 1.358 1.423.088-.064.165-.282.233-.626 1.344 1.303 3.242 4.586.488 5.889C19.367 9.343 16.568 7.938 13.466 6.885L13.466 6.885zM34.623 6.885c-3.986 2.055-6.305 3.718-7.574 5.134.65 2.607 4.043 2.726 5.283 2.653-.254-.119-.466-.26-.542-.479.312-.221 1.415-.023 2.186-.456-.296-.062-.434-.12-.573-.339.729-.232 1.514-.433 1.974-.817-.249.003-.481.055-.806-.17.652-.351 1.348-.629 1.889-1.166-.338-.009-.701-.004-.807-.129.598-.37 1.1-.78 1.518-1.23-.473.058-.671.009-.785-.075.451-.461 1.021-.85 1.293-1.421-.35.121-.67.168-.9-.011.152-.345.811-.55 1.188-1.357-.369.036-.76.081-.838 0 .172-.697.465-1.089.754-1.496-.789-.012-1.985.004-1.932-.063l.488-.499c-.771-.207-1.56.034-2.133.213-.258-.203.005-.459.318-.721-.654.087-1.248.237-1.782.445-.286-.258.186-.516.414-.774-1.013.191-1.44.461-1.867.731-.31-.295-.018-.548.19-.807-.763.283-1.156.648-1.57 1.008-.14-.189-.356-.327-.095-.783-.542.311-.951.68-1.252 1.092-.335-.215-.2-.506-.202-.775-.563.459-.92.946-1.358 1.423-.088-.064-.165-.282-.232-.626-1.345 1.303-3.243 4.586-.488 5.889C28.723 9.342 31.521 7.938 34.623 6.885L34.623 6.885z"/><g><path fill="#FF4081" d="M28.873 33.426c.014 2.433-2.113 4.414-4.75 4.428-2.638.012-4.788-1.948-4.801-4.381 0-.016 0-.031 0-.047-.014-2.433 2.112-4.414 4.75-4.428 2.638-.012 4.787 1.948 4.801 4.382C28.873 33.395 28.873 33.411 28.873 33.426zM21.333 20.846c1.979 1.296 2.335 4.234.797 6.563-1.539 2.329-4.391 3.165-6.37 1.868l0 0c-1.979-1.297-2.335-4.235-.797-6.563C16.502 20.385 19.355 19.549 21.333 20.846L21.333 20.846zM26.676 20.61c-1.98 1.295-2.337 4.235-.798 6.563 1.539 2.33 4.391 3.166 6.369 1.869l0 0c1.979-1.297 2.337-4.234.798-6.564C31.506 20.15 28.654 19.314 26.676 20.61L26.676 20.61zM11.443 22.966c2.136-.573.721 8.838-1.017 8.066C8.514 29.493 7.898 24.988 11.443 22.966zM36.135 22.848c-2.138-.572-.722 8.839 1.016 8.066C39.064 29.375 39.68 24.871 36.135 22.848zM28.875 15.839c3.687-.624 6.756 1.567 6.632 5.565C35.385 22.938 27.516 16.065 28.875 15.839zM18.687 15.72c-3.687-.621-6.755 1.57-6.631 5.567C12.177 22.821 20.045 15.949 18.687 15.72zM23.983 14.789c-2.2-.058-4.313 1.634-4.318 2.613-.006 1.19 1.741 2.412 4.333 2.442 2.648.019 4.337-.977 4.347-2.205C28.354 16.246 25.937 14.767 23.983 14.789L23.983 14.789zM24.118 39.221c1.919-.084 4.493.619 4.499 1.549.031.905-2.336 2.947-4.626 2.907-2.373.103-4.699-1.943-4.668-2.651C19.287 39.984 22.212 39.174 24.118 39.221zM17.031 33.703c1.366 1.646 1.988 4.539.849 5.39-1.079.652-3.698.384-5.56-2.29-1.255-2.245-1.094-4.527-.212-5.199C13.426 30.801 15.463 31.884 17.031 33.703L17.031 33.703zM30.932 33.183c-1.479 1.731-2.301 4.888-1.223 5.905 1.03.791 3.799.681 5.842-2.156 1.484-1.906.988-5.087.141-5.934C34.431 30.026 32.623 31.271 30.932 33.183L30.932 33.183z"/></g></svg>`

const IconRaspberryPIZero = `<svg version="1.1" id="Layer_1" width="80" height="80" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" x="0px" y="0px"
	 viewBox="0 0 496.8 496.8" style="enable-background:new 0 0 496.8 496.8;" xml:space="preserve">
<path style="fill:#EAC113;" d="M144,428.4c0,6.4-5.6,12-12,12l0,0c-6.4,0-12-5.6-12-12v-360c0-6.4,5.6-12,12-12l0,0
	c6.4,0,12,5.6,12,12V428.4z"/>
<path style="fill:#E8A615;" d="M132,56.4L132,56.4c6.4,0,12,4.8,12,11.2v361.6c0,6.4-5.6,11.2-12,11.2l0,0"/>
<path style="fill:#EAC113;" d="M64,428.4c0,6.4-5.6,12-12,12l0,0c-6.4,0-12-5.6-12-12v-360c0-6.4,5.6-12,12-12l0,0
	c6.4,0,12,5.6,12,12V428.4z"/>
<path style="fill:#E8A615;" d="M52,56.4L52,56.4c6.4,0,12,4.8,12,11.2v361.6c0,6.4-5.6,11.2-12,11.2l0,0"/>
<path style="fill:#EAC113;" d="M224,428.4c0,6.4-5.6,12-12,12l0,0c-6.4,0-12-5.6-12-12v-360c0-6.4,5.6-12,12-12l0,0
	c6.4,0,12,5.6,12,12V428.4z"/>
<path style="fill:#E8A615;" d="M212,56.4L212,56.4c6.4,0,12,4.8,12,11.2v361.6c0,6.4-5.6,11.2-12,11.2l0,0"/>
<path style="fill:#EAC113;" d="M304,428.4c0,6.4-5.6,12-12,12l0,0c-6.4,0-12-5.6-12-12v-360c0-6.4,5.6-12,12-12l0,0
	c6.4,0,12,5.6,12,12V428.4z"/>
<path style="fill:#E8A615;" d="M292,56.4L292,56.4c6.4,0,12,4.8,12,11.2v361.6c0,6.4-5.6,11.2-12,11.2l0,0"/>
<path style="fill:#EAC113;" d="M384,428.4c0,6.4-5.6,12-12,12l0,0c-6.4,0-12-5.6-12-12v-360c0-6.4,5.6-12,12-12l0,0
	c6.4,0,12,5.6,12,12V428.4z"/>
<path style="fill:#E8A615;" d="M372,56.4L372,56.4c6.4,0,12,4.8,12,11.2v361.6c0,6.4-5.6,11.2-12,11.2l0,0"/>
<path style="fill:#EAC113;" d="M464,428.4c0,6.4-5.6,12-12,12l0,0c-6.4,0-12-5.6-12-12v-360c0-6.4,5.6-12,12-12l0,0
	c6.4,0,12,5.6,12,12V428.4z"/>
<path style="fill:#E8A615;" d="M452,56.4L452,56.4c6.4,0,12,4.8,12,11.2v361.6c0,6.4-5.6,11.2-12,11.2l0,0"/>
<path style="fill:#62C106;" d="M496,366.8c0,13.6-11.2,25.6-25.6,25.6H25.6C11.2,392.4,0,381.2,0,366.8V130
	c0-14.4,11.2-25.6,25.6-25.6h445.6c13.6,0,25.6,11.2,25.6,25.6v236.8H496z"/>
<path style="fill:#56B502;" d="M0,130.8c0-13.6,11.2-26.4,25.6-26.4h445.6c13.6,0,25.6,13.6,25.6,28v235.2
	c0,13.6-11.2,25.6-25.6,25.6"/>
<path style="fill:#51A500;" d="M464,335.6c0,10.4-10.4,16.8-20.8,16.8H52.8c-10.4,0-20.8-5.6-20.8-16.8V161.2
	c0-11.2,10.4-16.8,20.8-16.8h390.4c10.4,0,20.8,5.6,20.8,16.8L464,335.6L464,335.6z"/>
<path style="fill:#4C9300;" d="M32,161.2c0-11.2,10.4-16.8,20.8-16.8h390.4c10.4,0,20.8,5.6,20.8,16.8v175.2
	c0,10.4-10.4,16.8-20.8,16.8"/>
<g>
	<polygon style="fill:#62C106;" points="168,320.4 104,320.4 104,176.4 72,176.4 72,168.4 112,168.4 112,312.4 160,312.4 160,256.4 
		232,256.4 232,264.4 168,264.4 	"/>
	<circle style="fill:#62C106;" cx="67.2" cy="174.8" r="11.2"/>
	<path style="fill:#62C106;" d="M244.8,264.4c0,6.4-4.8,11.2-11.2,11.2s-11.2-4.8-11.2-11.2s4.8-11.2,11.2-11.2
		S244.8,258,244.8,264.4z"/>
</g>
<polygon style="fill:#56B502;" points="376,160.4 432,160.4 432,304.4 472,304.4 472,312.4 424,312.4 424,168.4 384,168.4 
	384,224.4 312,224.4 312,216.4 376,216.4 "/>
<g>
	<circle style="fill:#62C106;" cx="307.2" cy="223.6" r="11.2"/>
	<polygon style="fill:#62C106;" points="224,160.4 280,160.4 280,304.4 312,304.4 312,312.4 272,312.4 272,168.4 232,168.4 
		232,224.4 152,224.4 152,216.4 224,216.4 	"/>
	<circle style="fill:#62C106;" cx="320.8" cy="312.4" r="11.2"/>
	<circle style="fill:#62C106;" cx="154.4" cy="223.6" r="11.2"/>
</g>
<g>
</g>
<g>
</g>
<g>
</g>
<g>
</g>
<g>
</g>
<g>
</g>
<g>
</g>
<g>
</g>
<g>
</g>
<g>
</g>
<g>
</g>
<g>
</g>
<g>
</g>
<g>
</g>
<g>
</g>
</svg>
`

const IconJetsonNano = `<svg version="1.1" id="Layer_1" width="80" height="80" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" x="0px" y="0px"
viewBox="0 0 53 56" style="enable-background:new 0 0 53 56;" xml:space="preserve">
<style type="text/css">
.st0{fill:#77B900;}
</style>
<title>Nvidia_logo</title>
<desc>Created with Sketch.</desc>
<path class="st0" d="M5.8,14.8c0,0,4.7-6.9,14-7.6V4.7C9.4,5.6,0.5,14.3,0.5,14.3s5.1,14.7,19.3,16v-2.7C9.3,26.3,5.8,14.8,5.8,14.8
z"/>
<path class="st0" d="M19.8,22.4v2.4c-7.9-1.4-10.1-9.6-10.1-9.6s3.8-4.2,10.1-4.9V13l0,0c-3.3-0.4-5.9,2.7-5.9,2.7
S15.3,20.9,19.8,22.4L19.8,22.4z"/>
<path class="st0" d="M19.8,0.1v4.6c0.3,0,0.6,0,0.9-0.1c11.8-0.4,19.4,9.6,19.4,9.6S31.3,25,22.1,25c-0.8,0-1.6-0.1-2.4-0.2v2.9
c0.6,0.1,1.3,0.1,2,0.1c8.5,0,14.7-4.4,20.7-9.5c1,0.8,5,2.7,5.9,3.6c-5.7,4.8-18.9,8.6-26.4,8.6c-0.7,0-1.4,0-2.1-0.1v4h32.4V0.1
H19.8z"/>
<path class="st0" d="M19.8,10.3V7.2c0.3,0,0.6,0,0.9,0c8.5-0.3,14,7.3,14,7.3s-6,8.3-12.4,8.3c-0.9,0-1.8-0.1-2.5-0.4V13
c3.3,0.4,4,1.9,5.9,5.1l4.4-3.7c0,0-3.2-4.2-8.6-4.2C20.9,10.2,20.3,10.3,19.8,10.3L19.8,10.3z"/>
<g>
<path d="M12,53.3H9.2l-5.5-8.9v8.9H1V39.7h2.8l5.5,9v-9H12V53.3z"/>
<path d="M22.3,50.5h-4.9l-0.9,2.8h-3l5.1-13.6h2.6l5.1,13.6h-3L22.3,50.5z M18.2,48.2h3.4l-1.7-5.1L18.2,48.2z"/>
<path d="M38.6,53.3h-2.8l-5.5-8.9v8.9h-2.8V39.7h2.8l5.5,9v-9h2.8V53.3z"/>
<path d="M52.2,46.8c0,1-0.1,2-0.4,2.8c-0.3,0.8-0.7,1.5-1.2,2.1c-0.5,0.6-1.1,1-1.8,1.3c-0.7,0.3-1.5,0.5-2.3,0.5
   c-0.9,0-1.6-0.2-2.3-0.5c-0.7-0.3-1.3-0.7-1.8-1.3c-0.5-0.6-0.9-1.3-1.2-2.1s-0.4-1.8-0.4-2.8v-0.6c0-1,0.1-2,0.4-2.8
   s0.7-1.5,1.2-2.1c0.5-0.6,1.1-1,1.8-1.3c0.7-0.3,1.5-0.5,2.3-0.5c0.9,0,1.6,0.2,2.3,0.5c0.7,0.3,1.3,0.8,1.8,1.3
   c0.5,0.6,0.9,1.3,1.2,2.1c0.3,0.8,0.4,1.8,0.4,2.8V46.8z M49.4,46.2c0-1.4-0.3-2.5-0.8-3.2c-0.5-0.7-1.2-1.1-2.2-1.1
   s-1.7,0.4-2.2,1.1c-0.5,0.7-0.8,1.8-0.8,3.2v0.6c0,0.7,0.1,1.3,0.2,1.9c0.1,0.5,0.3,1,0.6,1.4c0.3,0.4,0.6,0.7,0.9,0.8
   c0.4,0.2,0.8,0.3,1.3,0.3c0.9,0,1.7-0.4,2.2-1.1c0.5-0.7,0.8-1.8,0.8-3.3V46.2z"/>
</g>
</svg>`

const IconESP32 = `<svg version="1.1" id="Layer_1" width="70" height="70" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink"
	 viewBox="0 0 512 512" style="enable-background:new 0 0 512 512;right:20px" xml:space="preserve">
<g>
	<g>
		<g>
			<path d="M192.533,397.08h10.763v27.574c0,4.432,3.592,8.025,8.025,8.025s8.025-3.593,8.025-8.025V397.08h12.483v27.574
				c0,4.432,3.592,8.025,8.025,8.025s8.025-3.593,8.025-8.025V397.08h12.482v27.574c0,4.432,3.592,8.025,8.025,8.025
				s8.025-3.593,8.025-8.025V397.08h12.483v27.574c0,4.432,3.592,8.025,8.025,8.025s8.025-3.593,8.025-8.025V397.08h9.567
				c4.433,0,8.025-3.593,8.025-8.025v-20.173h19.795c13.274,0,24.074-10.799,24.074-24.074v-21.674h20.173
				c4.433,0,8.025-3.593,8.025-8.025v-9.224h29.113c4.433,0,8.025-3.593,8.025-8.025c0-4.432-3.592-8.025-8.025-8.025h-29.113
				v-12.483h29.113c4.433,0,8.025-3.593,8.025-8.025c0-4.432-3.592-8.025-8.025-8.025h-29.113v-12.482h29.113
				c4.433,0,8.025-3.593,8.025-8.025c0-4.432-3.592-8.025-8.025-8.025h-29.113v-12.483h29.113c4.433,0,8.025-3.593,8.025-8.025
				c0-4.432-3.592-8.025-8.025-8.025h-29.113v-11.105c0-4.432-3.592-8.025-8.025-8.025h-20.173v-17.914
				c0-13.275-10.8-24.074-24.074-24.074h-19.795v-22.053c0-4.432-3.592-8.025-8.025-8.025h-9.567V85.465
				c0-4.432-3.592-8.025-8.025-8.025s-8.025,3.593-8.025,8.025v27.574h-12.483V85.465c0-4.432-3.592-8.025-8.025-8.025
				s-8.025,3.593-8.025,8.025v27.574h-12.482V85.465c0-4.432-3.592-8.025-8.025-8.025s-8.025,3.593-8.025,8.025v27.574h-12.483
				V85.465c0-4.432-3.592-8.025-8.025-8.025s-8.025,3.593-8.025,8.025v27.574h-10.763c-4.433,0-8.025,3.593-8.025,8.025v22.053
				h-19.795c-13.274,0-24.075,10.799-24.075,24.074v17.914h-20.173c-4.433,0-8.025,3.593-8.025,8.025v11.105H84.526
				c-4.433,0-8.025,3.593-8.025,8.025c0,4.432,3.592,8.025,8.025,8.025h27.917v12.483H84.526c-4.433,0-8.025,3.593-8.025,8.025
				s3.592,8.025,8.025,8.025h27.917v12.482H84.526c-4.433,0-8.025,3.593-8.025,8.025c0,4.432,3.592,8.025,8.025,8.025h27.917v12.483
				H84.526c-4.433,0-8.025,3.593-8.025,8.025s3.592,8.025,8.025,8.025h27.917v9.224c0,4.432,3.592,8.025,8.025,8.025h20.173v21.676
				c0,13.275,10.8,24.074,24.074,24.074h19.795v20.173C184.508,393.487,188.1,397.08,192.533,397.08z M140.64,307.083h-12.149
				V201.156h12.149V307.083z M306.486,381.031H200.558v-12.148h105.928V381.031z M306.486,129.09v12.147H200.558V129.09H306.486z
				 M164.714,352.833c-4.425,0-8.025-3.599-8.025-8.025v-29.7V193.131v-25.938c0-4.425,3.599-8.025,8.025-8.025H342.33
				c4.425,0,8.025,3.599,8.025,8.025v25.938v24.609c0,4.432,3.592,8.025,8.025,8.025s8.025-3.593,8.025-8.025v-16.585h12.149
				v105.928h-12.149V243.42c0-4.432-3.592-8.025-8.025-8.025s-8.025,3.593-8.025,8.025v71.688v29.7c0,4.425-3.599,8.025-8.025,8.025
				h-27.819H192.533H164.714z"/>
			<path d="M334.305,328.759V183.242c0-4.432-3.592-8.025-8.025-8.025H180.763c-4.433,0-8.025,3.593-8.025,8.025v145.517
				c0,4.432,3.592,8.025,8.025,8.025H326.28C330.713,336.784,334.305,333.191,334.305,328.759z M318.255,320.734H188.788V191.267
				h129.467V320.734z"/>
			<path d="M76.799,51.915c0-13.078-10.639-23.718-23.717-23.718c-13.078,0-23.718,10.64-23.718,23.718s10.64,23.717,23.718,23.717
				C66.16,75.632,76.799,64.993,76.799,51.915z M53.082,59.583c-4.229,0-7.669-3.44-7.669-7.667c0-4.229,3.44-7.669,7.669-7.669
				c4.227,0,7.667,3.44,7.667,7.669C60.75,56.143,57.31,59.583,53.082,59.583z"/>
			<path d="M460.084,75.632c13.078,0,23.717-10.64,23.717-23.717c0-13.078-10.639-23.718-23.717-23.718
				c-13.078,0-23.718,10.64-23.718,23.718C436.366,64.993,447.006,75.632,460.084,75.632z M460.084,44.247
				c4.227,0,7.667,3.44,7.667,7.669s-3.44,7.667-7.667,7.667c-4.229,0-7.669-3.44-7.669-7.667
				C452.416,47.687,455.856,44.247,460.084,44.247z"/>
			<path d="M460.084,484.296c13.078,0,23.717-10.64,23.717-23.718c0-13.078-10.639-23.718-23.717-23.718
				c-13.078,0-23.718,10.64-23.718,23.718C436.366,473.656,447.006,484.296,460.084,484.296z M460.084,452.911
				c4.227,0,7.667,3.44,7.667,7.669c0,4.229-3.44,7.669-7.667,7.669c-4.229,0-7.669-3.44-7.669-7.669
				C452.416,456.349,455.856,452.911,460.084,452.911z"/>
			<path d="M463.428,0h-76.85c-4.433,0-8.025,3.593-8.025,8.025v18.926c0,2.655-2.159,4.815-4.815,4.815H161.504
				c-2.655,0-4.815-2.16-4.815-4.815V8.025c0-4.432-3.592-8.025-8.025-8.025H48.571C21.79,0,0.001,21.789,0.001,48.572v414.857
				C0.001,490.211,21.79,512,48.571,512h100.093c4.433,0,8.025-3.593,8.025-8.025v-18.926c0-2.655,2.16-4.815,4.815-4.815h23.031
				c4.433,0,8.025-3.593,8.025-8.025c0-4.432-3.592-8.025-8.025-8.025h-23.031c-11.504,0-20.865,9.36-20.865,20.865v10.901H48.571
				c-17.933,0-32.521-14.589-32.521-32.522V48.572c0-17.933,14.589-32.522,32.521-32.522h92.068v10.901
				c0,11.504,9.36,20.865,20.865,20.865h212.233c11.504,0,20.865-9.36,20.865-20.865V16.05h68.825
				c17.933,0,32.522,14.589,32.522,32.522v110.596h-20.173c-4.433,0-8.025,3.593-8.025,8.025v177.616
				c0,4.432,3.592,8.025,8.025,8.025h20.173v110.595c0,17.933-14.589,32.522-32.522,32.522h-68.825v-10.901
				c0-11.504-9.36-20.865-20.865-20.865H211.286c-4.433,0-8.025,3.593-8.025,8.025c0,4.432,3.592,8.025,8.025,8.025h162.453
				c2.656,0,4.815,2.16,4.815,4.815v18.926c0,4.432,3.592,8.025,8.025,8.025h76.85c26.782,0,48.572-21.789,48.572-48.572v-118.62
				V167.192V48.572C511.998,21.789,490.209,0,463.428,0z M495.949,336.784h-12.147V175.217h12.147V336.784z"/>
			<path d="M210.709,220.308h86.82c4.433,0,8.025-3.593,8.025-8.025c0-4.432-3.592-8.025-8.025-8.025h-86.82
				c-4.433,0-8.025,3.593-8.025,8.025C202.684,216.715,206.277,220.308,210.709,220.308z"/>
			<path d="M210.709,307.859h86.82c4.433,0,8.025-3.593,8.025-8.025s-3.592-8.025-8.025-8.025h-86.82
				c-4.433,0-8.025,3.593-8.025,8.025S206.277,307.859,210.709,307.859z"/>
			<path d="M210.709,278.675h86.82c4.433,0,8.025-3.593,8.025-8.025c0-4.432-3.592-8.025-8.025-8.025h-86.82
				c-4.433,0-8.025,3.593-8.025,8.025C202.684,275.083,206.277,278.675,210.709,278.675z"/>
			<path d="M210.709,249.492h86.82c4.433,0,8.025-3.593,8.025-8.025s-3.592-8.025-8.025-8.025h-86.82
				c-4.433,0-8.025,3.593-8.025,8.025S206.277,249.492,210.709,249.492z"/>
			<path d="M48.672,343.168h-0.022c-4.433,0-8.013,3.593-8.013,8.025s3.604,8.025,8.037,8.025c4.433,0,8.025-3.593,8.025-8.025
				S53.105,343.168,48.672,343.168z"/>
			<path d="M48.672,430.719h-0.022c-4.433,0-8.013,3.593-8.013,8.025s3.604,8.025,8.037,8.025c4.433,0,8.025-3.593,8.025-8.025
				C56.696,434.312,53.105,430.719,48.672,430.719z"/>
			<path d="M48.672,401.536h-0.022c-4.433,0-8.013,3.593-8.013,8.025c0,4.432,3.604,8.025,8.037,8.025
				c4.433,0,8.025-3.593,8.025-8.025C56.696,405.129,53.105,401.536,48.672,401.536z"/>
			<path d="M48.672,372.352h-0.022c-4.433,0-8.013,3.593-8.013,8.025c0,4.432,3.604,8.025,8.037,8.025
				c4.433,0,8.025-3.593,8.025-8.025C56.698,375.945,53.105,372.352,48.672,372.352z"/>
			<path d="M48.672,92.067h-0.022c-4.433,0-8.013,3.593-8.013,8.025c0,4.432,3.604,8.025,8.037,8.025
				c4.433,0,8.025-3.593,8.025-8.025C56.698,95.66,53.105,92.067,48.672,92.067z"/>
			<path d="M48.672,179.618h-0.022c-4.433,0-8.013,3.593-8.013,8.025c0,4.432,3.604,8.025,8.037,8.025
				c4.433,0,8.025-3.593,8.025-8.025C56.696,183.211,53.105,179.618,48.672,179.618z"/>
			<path d="M48.672,150.434h-0.022c-4.433,0-8.013,3.593-8.013,8.025s3.604,8.025,8.037,8.025c4.433,0,8.025-3.593,8.025-8.025
				C56.696,154.027,53.105,150.434,48.672,150.434z"/>
			<path d="M48.672,121.251h-0.022c-4.433,0-8.013,3.593-8.013,8.025c0,4.432,3.604,8.025,8.037,8.025
				c4.433,0,8.025-3.593,8.025-8.025C56.696,124.844,53.105,121.251,48.672,121.251z"/>
			<path d="M79.948,474.547c4.433,0,8.025-3.605,8.025-8.037s-3.592-8.025-8.025-8.025c-4.432,0-8.025,3.593-8.025,8.025v0.022
				C71.923,470.965,75.516,474.547,79.948,474.547z"/>
			<path d="M101.107,466.51v0.022c0,4.432,3.592,8.013,8.025,8.013c4.433,0,8.025-3.605,8.025-8.037
				c0-4.432-3.592-8.025-8.025-8.025C104.699,458.485,101.107,462.078,101.107,466.51z"/>
		</g>
	</g>
</g>
<g>
</g>
<g>
</g>
<g>
</g>
<g>
</g>
<g>
</g>
<g>
</g>
<g>
</g>
<g>
</g>
<g>
</g>
<g>
</g>
<g>
</g>
<g>
</g>
<g>
</g>
<g>
</g>
<g>
</g>
</svg>
`

const IconSparkFunBLE = `<svg version="1.1" id="Layer_1" width="80" height="60" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" x="0px" y="0px"
	 viewBox="0 0 501.6 501.6" style="enable-background:new 0 0 501.6 501.6;" xml:space="preserve">
<g>
	<path style="fill:#005EF4;" d="M218.8,404c-13.6,0-24-10.4-24-24v-91.2c0-13.6,10.4-24,24-24s24,10.4,24,24v90.4
		C242.8,392.8,232.4,404,218.8,404z"/>
	<path style="fill:#005EF4;" d="M218.8,140.8c-13.6,0-24-10.4-24-24V24c0-13.6,10.4-24,24-24s24,10.4,24,24v92.8
		C242.8,129.6,232.4,140.8,218.8,140.8z"/>
</g>
<path style="fill:#36A4FF;" d="M222,500.8c-6.4,0-12-2.4-16.8-7.2c-9.6-9.6-9.6-24.8,0-33.6l114.4-113.6L121.2,148
	c-9.6-9.6-9.6-24.8,0-33.6c9.6-9.6,24.8-9.6,33.6,0l232,232l-147.2,148C234.8,498.4,228.4,500.8,222,500.8z"/>
<path style="fill:#1085F9;" d="M155.6,114.4l232,232l-148,148c-4.8,4.8-11.2,7.2-16.8,7.2s-12-2.4-16.8-7.2
	c-9.6-9.6-9.6-24.8,0-33.6l113.6-114.4L121.2,148"/>
<rect x="252.395" y="270.655" transform="matrix(0.7071 0.7071 -0.7071 0.7071 296.8044 -127.2401)" style="fill:#005EF4;" width="99.199" height="48"/>
<path style="fill:#36A4FF;" d="M138,393.6c-6.4,0-12-2.4-16.8-7.2c-9.6-9.6-9.6-24.8,0-33.6l198.4-198.4L205.2,40.8
	c-9.6-9.6-9.6-24.8,0-33.6c9.6-9.6,24.8-9.6,33.6,0l148.8,147.2l-232,232C150.8,391.2,144.4,393.6,138,393.6z"/>
<path style="fill:#1085F9;" d="M121.2,352.8l198.4-198.4L205.2,40.8c-9.6-9.6-9.6-24.8,0-33.6c9.6-9.6,24.8-9.6,33.6,0l148.8,147.2
	l-232,232"/>
<g>
</g>
<g>
</g>
<g>
</g>
<g>
</g>
<g>
</g>
<g>
</g>
<g>
</g>
<g>
</g>
<g>
</g>
<g>
</g>
<g>
</g>
<g>
</g>
<g>
</g>
<g>
</g>
<g>
</g>
</svg>
`

var IconMap = map[int]template.HTML{
	usage.UnknownDevice:    "",
	usage.RaspberryPI_3B:   IconRaspberryPI,
	usage.RaspberryPI_4B:   IconRaspberryPI,
	usage.RaspberryPI_Zero: IconRaspberryPIZero,
	usage.JetsonNano:       IconJetsonNano,
	usage.SparkfunBLE:      IconSparkFunBLE,
	usage.ESP32CAM:         IconESP32,
}

/*var IconMap = map[int]template.HTML{
	usage.UnknownDevice:    "",
	usage.RaspberryPI_3B:   "ion-ios-pulse",
	usage.RaspberryPI_4B:   "ion-ios-pulse",
	usage.RaspberryPI_Zero: "ion-camera",
	usage.JetsonNano:       "ion-cloud",
	usage.SparkfunBLE:      "ion-bluetooth",
	usage.ESP32CAM:         "ion-android-wifi",
}
*/
