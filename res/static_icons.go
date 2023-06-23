package res

import "fyne.io/fyne/v2"

var Dumbell = &fyne.StaticResource{
	StaticName: "dumbell",
	StaticContent: []byte(
		`<?xml version="1.0" encoding="utf-8"?>
<svg width="800px" height="800px" viewBox="2 2 22 22" fill="none" xmlns="http://www.w3.org/2000/svg">
<path d="M10.81 20.38C10.2076 20.3793 9.62953 20.1423 9.2 19.72L4.2 14.72C3.98805 14.5099 3.81982 14.2599 3.70501 13.9844C3.5902 13.7089 3.53109 13.4134 3.53109 13.115C3.53109 12.8166 3.5902 12.5211 3.70501 12.2456C3.81982 11.9701 3.98805 11.7201 4.2 11.51L4.43 11.27C4.85868 10.8462 5.43718 10.6085 6.04 10.6085C6.64281 10.6085 7.22132 10.8462 7.65 11.27L12.65 16.27C12.8619 16.4801 13.0302 16.7301 13.145 17.0056C13.2598 17.2811 13.3189 17.5766 13.3189 17.875C13.3189 18.1734 13.2598 18.4689 13.145 18.7444C13.0302 19.0199 12.8619 19.2699 12.65 19.48L12.42 19.72C11.9905 20.1423 11.4124 20.3793 10.81 20.38ZM6.05 12.11C5.94986 12.1088 5.85049 12.1276 5.75774 12.1654C5.665 12.2032 5.58076 12.2591 5.51 12.33L5.27 12.57C5.19834 12.6415 5.14149 12.7265 5.1027 12.82C5.06391 12.9135 5.04394 13.0137 5.04394 13.115C5.04394 13.2162 5.06391 13.3165 5.1027 13.41C5.14149 13.5035 5.19834 13.5885 5.27 13.66L10.27 18.66C10.4161 18.8053 10.6139 18.8869 10.82 18.8869C11.0261 18.8869 11.2238 18.8053 11.37 18.66L11.6 18.42C11.6734 18.3486 11.7317 18.263 11.7712 18.1685C11.8107 18.074 11.8307 17.9724 11.83 17.87C11.8305 17.7691 11.8104 17.6691 11.7708 17.5762C11.7313 17.4833 11.6731 17.3996 11.6 17.33L6.6 12.33C6.52739 12.2588 6.4414 12.2027 6.347 12.165C6.2526 12.1272 6.15166 12.1085 6.05 12.11Z" fill="#000000"/>
<path d="M17.92 13.37C17.3164 13.3656 16.7385 13.1251 16.31 12.7L11.31 7.7C10.8838 7.2726 10.6444 6.69362 10.6444 6.09C10.6444 5.48639 10.8838 4.9074 11.31 4.48L11.55 4.24C11.977 3.81681 12.5538 3.57938 13.155 3.57938C13.7562 3.57938 14.333 3.81681 14.76 4.24L19.76 9.24C20.1832 9.66699 20.4206 10.2438 20.4206 10.845C20.4206 11.4462 20.1832 12.023 19.76 12.45L19.52 12.69C19.0956 13.1164 18.5215 13.3603 17.92 13.37ZM13.16 5.09C13.059 5.08888 12.9588 5.10875 12.8659 5.14834C12.7729 5.18793 12.6892 5.24638 12.62 5.32L12.38 5.55C12.2347 5.69615 12.1531 5.89389 12.1531 6.1C12.1531 6.30612 12.2347 6.50385 12.38 6.65L17.38 11.65C17.5261 11.7953 17.7239 11.8769 17.93 11.8769C18.1361 11.8769 18.3338 11.7953 18.48 11.65L18.71 11.41C18.7831 11.3404 18.8413 11.2566 18.8808 11.1638C18.9204 11.0709 18.9405 10.9709 18.94 10.87C18.9407 10.7676 18.9207 10.666 18.8812 10.5715C18.8417 10.477 18.7834 10.3914 18.71 10.32L13.71 5.32C13.6382 5.24703 13.5526 5.18911 13.4582 5.14961C13.3637 5.11012 13.2624 5.08986 13.16 5.09Z" fill="#000000"/>
<path d="M19.55 10.21C18.9476 10.2093 18.3695 9.97235 17.94 9.55L14.53 6.13C14.3181 5.91989 14.1498 5.66988 14.035 5.39441C13.9202 5.11893 13.8611 4.82344 13.8611 4.525C13.8611 4.22656 13.9202 3.93107 14.035 3.6556C14.1498 3.38012 14.3181 3.13011 14.53 2.92L14.76 2.68C15.1874 2.25376 15.7664 2.0144 16.37 2.0144C16.9736 2.0144 17.5526 2.25376 17.98 2.68L21.4 6.09C21.8238 6.51868 22.0615 7.09719 22.0615 7.7C22.0615 8.30282 21.8238 8.88132 21.4 9.31L21.16 9.55C20.7305 9.97235 20.1524 10.2093 19.55 10.21ZM16.37 3.51C16.1639 3.51294 15.9669 3.59534 15.82 3.74L15.59 4C15.5183 4.07152 15.4615 4.15647 15.4227 4.24999C15.3839 4.34351 15.3639 4.44376 15.3639 4.545C15.3639 4.64624 15.3839 4.74649 15.4227 4.84001C15.4615 4.93353 15.5183 5.01848 15.59 5.09L19 8.49C19.1462 8.63534 19.3439 8.71692 19.55 8.71692C19.7561 8.71692 19.9538 8.63534 20.1 8.49L20.33 8.25C20.4032 8.17831 20.4613 8.09273 20.501 7.99829C20.5407 7.90385 20.5611 7.80244 20.5611 7.7C20.5611 7.59756 20.5407 7.49615 20.501 7.40171C20.4613 7.30727 20.4032 7.22169 20.33 7.15L16.92 3.74C16.7744 3.59353 16.5766 3.51082 16.37 3.51Z" fill="#000000"/>
<path d="M7.63 22C7.32648 22.0025 7.02561 21.9435 6.74549 21.8266C6.46536 21.7098 6.21178 21.5374 6 21.32L2.61 17.91C2.39775 17.699 2.2293 17.4482 2.11436 17.1719C1.99942 16.8956 1.94025 16.5993 1.94025 16.3C1.94025 16.0007 1.99942 15.7044 2.11436 15.4281C2.2293 15.1518 2.39775 14.901 2.61 14.69L2.84 14.45C3.26868 14.0262 3.84718 13.7885 4.45 13.7885C5.05281 13.7885 5.63131 14.0262 6.06 14.45L9.47 17.87C9.68194 18.0801 9.85018 18.3301 9.96498 18.6056C10.0798 18.8811 10.1389 19.1766 10.1389 19.475C10.1389 19.7734 10.0798 20.0689 9.96498 20.3444C9.85018 20.6199 9.68194 20.8699 9.47 21.08L9.24 21.32C9.03022 21.5346 8.77981 21.7052 8.50339 21.8219C8.22697 21.9387 7.93006 21.9992 7.63 22ZM4.45 15.3C4.34834 15.2985 4.24739 15.3172 4.15299 15.355C4.05859 15.3928 3.97261 15.4488 3.9 15.52L3.67 15.76C3.59683 15.8317 3.53869 15.9173 3.49901 16.0117C3.45933 16.1062 3.43888 16.2076 3.43888 16.31C3.43888 16.4124 3.45933 16.5139 3.49901 16.6083C3.53869 16.7027 3.59683 16.7883 3.67 16.86L7.08 20.27C7.15152 20.3417 7.23647 20.3985 7.32999 20.4373C7.4235 20.4761 7.52375 20.4961 7.625 20.4961C7.72624 20.4961 7.82649 20.4761 7.92001 20.4373C8.01353 20.3985 8.09848 20.3417 8.17 20.27L8.41 20C8.48166 19.9285 8.53851 19.8435 8.5773 19.75C8.61609 19.6565 8.63605 19.5562 8.63605 19.455C8.63605 19.3538 8.61609 19.2535 8.5773 19.16C8.53851 19.0665 8.48166 18.9815 8.41 18.91L5 15.51C4.92738 15.4388 4.8414 15.3828 4.747 15.345C4.6526 15.3072 4.55166 15.2885 4.45 15.29V15.3Z" fill="#000000"/>
<path d="M10.78 16.2C10.6812 16.2022 10.583 16.1838 10.4918 16.1459C10.4005 16.108 10.3182 16.0515 10.25 15.98L8 13.77C7.92924 13.7011 7.87301 13.6186 7.83461 13.5276C7.79621 13.4366 7.77643 13.3388 7.77643 13.24C7.77643 13.1412 7.79621 13.0434 7.83461 12.9524C7.87301 12.8614 7.92924 12.7789 8 12.71L12.71 8C12.7782 7.92846 12.8605 7.87195 12.9518 7.83407C13.043 7.79618 13.1412 7.77777 13.24 7.78C13.3388 7.77777 13.437 7.79618 13.5282 7.83407C13.6195 7.87195 13.7018 7.92846 13.77 8L16 10.25C16.0708 10.3189 16.127 10.4014 16.1654 10.4924C16.2038 10.5834 16.2236 10.6812 16.2236 10.78C16.2236 10.8788 16.2038 10.9766 16.1654 11.0676C16.127 11.1586 16.0708 11.2411 16 11.31L11.31 16C11.2401 16.0679 11.157 16.1207 11.0658 16.1551C10.9746 16.1896 10.8773 16.2048 10.78 16.2ZM9.63 13.2L10.78 14.35L14.39 10.74L13.24 9.63L9.63 13.2Z" fill="#000000"/>
</svg>`,
	),
}

var City = &fyne.StaticResource{
	StaticName: "city",
	StaticContent: []byte(
		`<?xml version="1.0" encoding="utf-8"?><svg fill="#000000" width="800px" height="800px" viewBox="2 2 22 22" xmlns="http://www.w3.org/2000/svg"><path d="M13,9a1,1,0,0,0-1-1H3A1,1,0,0,0,2,9V22H13ZM6,20H4V18H6Zm0-4H4V14H6Zm0-4H4V10H6Zm5,8H8V18h3Zm0-4H8V14h3Zm0-4H8V10h3Zm3.5-6H6V3A1,1,0,0,1,7,2H17a1,1,0,0,1,1,1v7H15V6.5A.5.5,0,0,0,14.5,6ZM22,13v9H19.5V18h-2v4H15V13a1,1,0,0,1,1-1h5A1,1,0,0,1,22,13Z"/></svg>`,
	),
}

var Team = &fyne.StaticResource{
	StaticName: "team",
	StaticContent: []byte(
		`<?xml version="1.0" encoding="utf-8"?>
        <!-- Uploaded to: SVG Repo, www.svgrepo.com, Generator: SVG Repo Mixer Tools -->
        <svg fill="#000000" width="800px" height="800px" viewBox="-3 0 32 32" version="1.1" xmlns="http://www.w3.org/2000/svg">
        <title>group</title>
        <path d="M20.906 20.75c1.313 0.719 2.063 2 1.969 3.281-0.063 0.781-0.094 0.813-1.094 0.938-0.625 0.094-4.563 0.125-8.625 0.125-4.594 0-9.406-0.094-9.75-0.188-1.375-0.344-0.625-2.844 1.188-4.031 1.406-0.906 4.281-2.281 5.063-2.438 1.063-0.219 1.188-0.875 0-3-0.281-0.469-0.594-1.906-0.625-3.406-0.031-2.438 0.438-4.094 2.563-4.906 0.438-0.156 0.875-0.219 1.281-0.219 1.406 0 2.719 0.781 3.25 1.938 0.781 1.531 0.469 5.625-0.344 7.094-0.938 1.656-0.844 2.188 0.188 2.469 0.688 0.188 2.813 1.188 4.938 2.344zM3.906 19.813c-0.5 0.344-0.969 0.781-1.344 1.219-1.188 0-2.094-0.031-2.188-0.063-0.781-0.188-0.344-1.625 0.688-2.25 0.781-0.5 2.375-1.281 2.813-1.375 0.563-0.125 0.688-0.469 0-1.656-0.156-0.25-0.344-1.063-0.344-1.906-0.031-1.375 0.25-2.313 1.438-2.719 1-0.375 2.125 0.094 2.531 0.938 0.406 0.875 0.188 3.125-0.25 3.938-0.5 0.969-0.406 1.219 0.156 1.375 0.125 0.031 0.375 0.156 0.719 0.313-1.375 0.563-3.25 1.594-4.219 2.188zM24.469 18.625c0.75 0.406 1.156 1.094 1.094 1.813-0.031 0.438-0.031 0.469-0.594 0.531-0.156 0.031-0.875 0.063-1.813 0.063-0.406-0.531-0.969-1.031-1.656-1.375-1.281-0.75-2.844-1.563-4-2.063 0.313-0.125 0.594-0.219 0.719-0.25 0.594-0.125 0.688-0.469 0-1.656-0.125-0.25-0.344-1.063-0.344-1.906-0.031-1.375 0.219-2.313 1.406-2.719 1.031-0.375 2.156 0.094 2.531 0.938 0.406 0.875 0.25 3.125-0.188 3.938-0.5 0.969-0.438 1.219 0.094 1.375 0.375 0.125 1.563 0.688 2.75 1.313z"></path>
        </svg>`,
	),
}

var Money = &fyne.StaticResource{
	StaticName: "money",
	StaticContent: []byte(
		`<?xml version="1.0" encoding="utf-8"?><!-- Uploaded to: SVG Repo, www.svgrepo.com, Generator: SVG Repo Mixer Tools -->
        <svg fill="#000000" width="800px" height="800px" viewBox="0 0 32 32" xmlns="http://www.w3.org/2000/svg">
            <path d="M31,7H1A1,1,0,0,0,0,8V24a1,1,0,0,0,1,1H31a1,1,0,0,0,1-1V8A1,1,0,0,0,31,7ZM25.09,23H6.91A6,6,0,0,0,2,18.09V13.91A6,6,0,0,0,6.91,9H25.09A6,6,0,0,0,30,13.91v4.18A6,6,0,0,0,25.09,23ZM30,11.86A4,4,0,0,1,27.14,9H30ZM4.86,9A4,4,0,0,1,2,11.86V9ZM2,20.14A4,4,0,0,1,4.86,23H2ZM27.14,23A4,4,0,0,1,30,20.14V23Z"/>
            <path d="M7.51.71a1,1,0,0,0-.76-.1,1,1,0,0,0-.61.46l-2,3.43a1,1,0,0,0,1.74,1L7.38,2.94l5.07,2.93a1,1,0,0,0,1-1.74Z"/>
            <path d="M24.49,31.29a1,1,0,0,0,.5.14.78.78,0,0,0,.26,0,1,1,0,0,0,.61-.46l2-3.43a1,1,0,1,0-1.74-1l-1.48,2.56-5.07-2.93a1,1,0,0,0-1,1.74Z"/>
            <path d="M16,10a6,6,0,1,0,6,6A6,6,0,0,0,16,10Zm0,10a4,4,0,1,1,4-4A4,4,0,0,1,16,20Z"/>
        </svg>`,
	),
}

var SadFace = &fyne.StaticResource{
	StaticName: "sad_face",
	StaticContent: []byte(
		`<?xml version="1.0" encoding="utf-8"?>
        <svg version="1.1" id="_x32_" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" 
             width="800px" height="800px" viewBox="0 0 512 512"  xml:space="preserve">
        <g>
            <rect x="186.063" y="342.594" class="st0" width="139.875" height="32.313"/>
            <path class="st0" d="M256.016,0C114.625,0,0,114.625,0,256s114.625,256,256.016,256C397.375,512,512,397.375,512,256
                S397.375,0,256.016,0z M256.016,450.031c-107,0-194.047-87.063-194.047-194.031c0-37.766,11.031-72.953,29.781-102.781h328.5
                C439,183.047,450,218.234,450,256C450,362.969,362.969,450.031,256.016,450.031z"/>
            <path class="st0" d="M290.719,240.813c5.344,11.313,13.188,21.031,23,28.031c9.781,7,21.656,11.203,34.313,11.188
                c12.688,0.016,24.563-4.188,34.344-11.188c9.813-7,17.672-16.719,23-28.031l-23.656-11.141c-3.563,7.609-8.688,13.734-14.5,17.875
                c-5.844,4.141-12.25,6.328-19.188,6.344c-6.906-0.016-13.313-2.203-19.156-6.344c-5.813-4.141-10.922-10.297-14.5-17.891
                L290.719,240.813z"/>
            <path class="st0" d="M106.625,240.813c5.344,11.313,13.172,21.031,22.984,28.031c9.781,7,21.672,11.203,34.328,11.188
                c12.656,0.016,24.547-4.188,34.328-11.188c9.828-7,17.672-16.719,23.016-28.031l-23.656-11.141
                c-3.578,7.578-8.703,13.734-14.5,17.875c-5.859,4.141-12.281,6.328-19.188,6.344c-6.906-0.016-13.313-2.203-19.156-6.344
                c-5.813-4.141-10.938-10.297-14.531-17.891L106.625,240.813z"/>
        </g>
        </svg>`,
	),
}

var MehFace = &fyne.StaticResource{
	StaticName: "meh_face",
	StaticContent: []byte(
		`<?xml version="1.0" encoding="utf-8"?>
        <svg version="1.1" id="_x32_" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" 
             width="800px" height="800px" viewBox="0 0 512 512"  xml:space="preserve">
        <g>
            <path class="st0" d="M256.016,0C114.625,0,0,114.625,0,256s114.625,256,256.016,256C397.375,512,512,397.375,512,256
                S397.375,0,256.016,0z M256.016,450.031c-106.984,0-194.047-87.063-194.047-194.031c0-37.766,11.031-72.953,29.797-102.781H420.25
                C439,183.047,450,218.234,450,256C450,362.969,362.969,450.031,256.016,450.031z"/>
            <path class="st0" d="M164.094,214.547c-19,0-34.406,15.422-34.406,34.406c0,19.047,15.406,34.453,34.406,34.453
                c19.031,0,34.438-15.406,34.438-34.453C198.531,229.969,183.125,214.547,164.094,214.547z"/>
            <path class="st0" d="M350.313,214.547c-19.031,0-34.469,15.422-34.469,34.406c0,19.047,15.438,34.453,34.469,34.453
                c19,0,34.406-15.406,34.406-34.453C384.719,229.969,369.313,214.547,350.313,214.547z"/>
            <rect x="188.875" y="342.594" class="st0" width="139.875" height="32.313"/>
        </g>
        </svg>`,
	),
}

var HappyFace = &fyne.StaticResource{
	StaticName: "happy_face",
	StaticContent: []byte(
		`<?xml version="1.0" encoding="utf-8"?>
        <svg version="1.1" id="_x32_" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" 
             width="800px" height="800px" viewBox="0 0 512 512"  xml:space="preserve">
        <g>
            <path class="st0" d="M256.016,0C114.625,0,0,114.625,0,256s114.625,256,256.016,256C397.375,512,512,397.375,512,256
                S397.375,0,256.016,0z M256.016,450.031c-107,0-194.047-87.063-194.047-194.031c0-37.766,11.031-72.953,29.781-102.781h328.5
                C439,183.047,450,218.234,450,256C450,362.969,362.969,450.031,256.016,450.031z"/>
            <path class="st0" d="M221.266,261.375c-5.328-11.313-13.172-21.031-22.984-28.031c-9.781-6.984-21.672-11.188-34.344-11.188
                c-12.656,0-24.547,4.203-34.344,11.188c-9.813,7-17.656,16.719-22.984,28.063l23.656,11.125c3.578-7.594,8.688-13.75,14.5-17.875
                c5.859-4.156,12.266-6.344,19.172-6.344c6.922,0,13.328,2.188,19.156,6.344c5.813,4.125,10.938,10.281,14.531,17.875
                L221.266,261.375z"/>
            <path class="st0" d="M405.375,261.375c-5.344-11.313-13.188-21.031-23-28.031c-9.781-6.984-21.656-11.188-34.313-11.188
                s-24.563,4.203-34.344,11.188c-9.813,7-17.656,16.719-23,28.063l23.656,11.125c3.563-7.594,8.688-13.75,14.5-17.875
                c5.844-4.156,12.266-6.344,19.188-6.344c6.906,0,13.313,2.188,19.141,6.344c5.828,4.125,10.953,10.281,14.531,17.875
                L405.375,261.375z"/>
            <path class="st0" d="M167.781,344.844c17.234,30.844,50.281,51.813,88.203,51.813c37.953,0,71.016-20.969,88.234-51.828
                l-26.609-14.875c-12.109,21.609-35.109,36.203-61.625,36.203c-26.484,0-49.469-14.594-61.594-36.219L167.781,344.844z"/>
        </g>
        </svg>`,
	),
}

var Transfers = &fyne.StaticResource{
	StaticName: "transfers",
	StaticContent: []byte(
		`<svg viewBox="2 2 22 22" width="800px" height="800px" xmlns="http://www.w3.org/2000/svg" fill="#ffffff"><g id="SVGRepo_bgCarrier" stroke-width="0"></g><g id="SVGRepo_tracerCarrier" stroke-linecap="round" stroke-linejoin="round"></g><g id="SVGRepo_iconCarrier"> <path stroke="#ffffff" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 17h12M4 17l3.5-3.5M4 17l3.5 3.5M7 7h13m0 0l-3.5-3.5M20 7l-3.5 3.5"></path> </g></svg>`,
	),
}

var Contract = &fyne.StaticResource{
	StaticName: "contract",
	StaticContent: []byte(
		`<svg viewBox="0 0 1024 1024" class="icon" version="1.1" xmlns="http://www.w3.org/2000/svg" fill="#ffffff"><g id="SVGRepo_bgCarrier" stroke-width="0"></g><g id="SVGRepo_tracerCarrier" stroke-linecap="round" stroke-linejoin="round"></g><g id="SVGRepo_iconCarrier"><path d="M182.52 146.2h585.14v256h73.15V73.06H109.38v877.71h256v-73.14H182.52z" fill="#ffffff"></path><path d="M255.67 219.34h438.86v73.14H255.67zM255.67 365.63h365.71v73.14H255.67zM255.67 511.91H475.1v73.14H255.67zM775.22 458.24L439.04 794.42l-0.52 154.64 155.68 0.52L930.38 613.4 775.22 458.24z m51.72 155.16l-25.43 25.43-51.73-51.72 25.44-25.44 51.72 51.73z m-77.14 77.15L620.58 819.77l-51.72-51.72 129.22-129.22 51.72 51.72zM511.91 876.16l0.17-51.34 5.06-5.06 51.72 51.72-4.85 4.85-52.1-0.17z" fill="#ffffff"></path></g></svg>`,
	),
}

var EmailRead = &fyne.StaticResource{
	StaticName: "email_read",
	StaticContent: []byte(
		`<?xml version="1.0" encoding="utf-8"?>
        <svg width="800px" height="800px" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
        <g id="Communication / Mail_Open">
        <path id="Vector" d="M4 10L10.1076 14.6122L10.1097 14.6139C10.7878 15.1112 11.1271 15.36 11.4988 15.4561C11.8272 15.5411 12.1725 15.5411 12.501 15.4561C12.8729 15.3599 13.2132 15.1104 13.8926 14.6122L20 10M19.8002 9.03944L14.2012 4.55657C13.506 3.99995 13.1581 3.72174 12.7715 3.61126C12.4304 3.51378 12.0692 3.50861 11.7255 3.59661C11.336 3.69634 10.9809 3.96473 10.2705 4.50188L4.26953 9.03967C3.8038 9.39184 3.57123 9.56804 3.40332 9.7906C3.2546 9.98772 3.14377 10.2107 3.07624 10.4482C3 10.7163 3 11.0083 3 11.5922V17.8001C3 18.9202 3 19.4805 3.21799 19.9083C3.40973 20.2847 3.71547 20.5904 4.0918 20.7822C4.5192 20.9999 5.07899 20.9999 6.19691 20.9999H17.8031C18.921 20.9999 19.48 20.9999 19.9074 20.7822C20.2837 20.5904 20.5905 20.2844 20.7822 19.9081C21 19.4807 21 18.9214 21 17.8035V11.5265C21 10.9693 21 10.689 20.9287 10.4301C20.8651 10.1992 20.7595 9.98161 20.619 9.78768C20.4604 9.56876 20.2409 9.39227 19.8002 9.03944Z" stroke="#333333" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>
        </g>
        </svg>`,
	),
}

var Unknown = &fyne.StaticResource{
	StaticName: "unknown",
	StaticContent: []byte(
		`<svg viewBox="0 0 24 24" width="800px" height="800px" xmlns="http://www.w3.org/2000/svg" fill="#ffffff"><g id="SVGRepo_bgCarrier" stroke-width="0"></g><g id="SVGRepo_tracerCarrier" stroke-linecap="round" stroke-linejoin="round"></g><g id="SVGRepo_iconCarrier"> <path fill="none" stroke="#ffffff" stroke-width="2" d="M2,3.99079514 C2,2.89130934 2.89821238,2 3.99079514,2 L20.0092049,2 C21.1086907,2 22,2.89821238 22,3.99079514 L22,20.0092049 C22,21.1086907 21.1017876,22 20.0092049,22 L3.99079514,22 C2.89130934,22 2,21.1017876 2,20.0092049 L2,3.99079514 Z M12,15 L12,14 C12,13 12,12.5 13,12 C14,11.5 15,11 15,9.5 C15,8.5 14,7 12,7 C10,7 9,8.26413718 9,10 M12,16 L12,18"></path> </g></svg>`,
	),
}
