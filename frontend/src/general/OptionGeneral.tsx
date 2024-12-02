import { Events, WML } from "@wailsio/runtime";
import { useEffect } from "react";

function setTime(data: any) {
    throw new Error("Function not implemented.");
}
const General: React.FC = (props) => {

    useEffect(() => {
        Events.On('time', (timeValue: any) => {
            setTime(timeValue.data);
        });
        // Reload WML so it picks up the wml tags
        WML.Reload();
    }, []);

    return (
        <div className="w-full h-full">fuck general content1a12123123---123123i</div>
    )

}

export default General

