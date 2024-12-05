import { Events, WML } from "@wailsio/runtime";
import { Card, Input, Radio, Typography } from "antd";
import { useEffect, useState } from "react";
import { StartupService } from "../../bindings/mousk/service";

function setTime(data: any) {
    throw new Error("Function not implemented.");
}
const General: React.FC = (props) => {
    const { Title, Paragraph } = Typography;
    useEffect(() => {
        Events.On('time', (timeValue: any) => {
            setTime(timeValue.data);
        });

        const loadData = async () => {
            setAutoStart(await StartupService.StartupState())
        };
        loadData();
        // Reload WML so it picks up the wml tags
        WML.Reload();
    }, []);

    const [autoStart, setAutoStart] = useState(false);


    const handleAutoStartChange = (e) => {
        console.log("auto start", e.target.value)
        StartupService.Startup(e.target.value)
        setAutoStart(e.target.value);
    };
    return (
        // <div className=" bg-slate-100 min-h-screen flex items-center justify-center">
        <Card className="w-auto h-full text-center rounded-lg shadow-lg p-8">
            <Title level={2}>Mousk.</Title>
            {/* <Paragraph type="secondary">作者：zui5</Paragraph> */}

            <div className="mt-6 text-left">
                <Paragraph strong>是否随系统启动</Paragraph>
                <Radio.Group onChange={handleAutoStartChange} value={autoStart} className="block mt-2">
                    <Radio value={true}>是</Radio>
                    <Radio value={false} className="ml-6">否</Radio>
                </Radio.Group>
            </div>

            <div className="mt-6 text-left">
                <Paragraph strong>其他选项</Paragraph>
                <Input placeholder="留白" className="mt-2" />
            </div>
        </Card>
        // </div>
    )

}

export default General

