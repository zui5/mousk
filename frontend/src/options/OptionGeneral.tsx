import { Events, WML } from "@wailsio/runtime";
import { useEffect } from "react";

function General() {

    useEffect(() => {
        Events.On('time', (timeValue: any) => {
            setTime(timeValue.data);
        });
        // Reload WML so it picks up the wml tags
        WML.Reload();
    }, []);

    return (
        <div>fuck general content</div>
    )

}

export default General

function setTime(data: any) {
    throw new Error("Function not implemented.");
}
