import { Browser } from '@wailsio/runtime';
const About: React.FC = (props) => {


    function showAboutInfo() {
        Browser.OpenURL("https://github.com/zui5/mousk");
    }

    return (
        <div className="w-full h-full flex items-center justify-center">
            <div className="bg-white p-8 rounded-lg shadow-lg max-w-md text-center">
                <h1 className="text-2xl font-bold mb-4">About Us</h1>
                <p className="text-gray-700 mb-6">
                    Keyboard Control Mouse Pointer Project<br></br>
                    @AUTHOR: zui5
                    {/* <a href="https://github.com/zui5/mousk">github</a> */}
                </p>
                <button onClick={showAboutInfo} className="bg-blue-500 text-white px-4 py-2 rounded hover:bg-blue-700">
                GOTO Github
                </button>
            </div>
        </div>


    )
}

export default About