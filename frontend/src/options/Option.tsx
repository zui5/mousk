
import { Events, WML } from "@wailsio/runtime";
import { Layout, Menu, MenuTheme, Switch, type MenuProps } from 'antd';
import { Content } from "antd/es/layout/layout";
import Sider from "antd/es/layout/Sider";
import React, { useEffect, useState } from 'react';
import About from "../about/About";
import Helper from "../about/Helper";
import General from "../general/OptionGeneral";
import Keymap from "../keymap/Keymap";

type MenuItem = Required<MenuProps>['items'][number];

const items: MenuItem[] = [
  { key: 'General', label: 'General' },
  { key: 'Keymap', label: 'Keymap' },
  { key: 'Helper', label: 'Helper' },
  { key: 'About', label: 'About' },
];

const contentsMap: { [key: string]: React.FC } = {
  "General": General,
  "Keymap": Keymap,
  "Helper": Helper,
  "About": About,
};

function GetCurrentComponent(key: string): React.FC {
  return contentsMap[key] || (() => <div>not complete yet!</div>);
}

const Option: React.FC = () => {
  const [, setTime] = useState<string>('Listening for Time event...');
  const [current, setCurrent] = useState('General');
  const [collapsed, setCollapsed] = useState(false);
  // const {
  //   token: { colorBgContainer, borderRadiusLG },
  // } = theme.useToken();

  const onClick: MenuProps['onClick'] = (e) => {
    console.log('click ', e);
    setCurrent(e.key);
  };

  useEffect(() => {
    Events.On('time', (timeValue: any) => {
      setTime(timeValue.data);
    });
    WML.Reload();
  }, []);

  const CurrentComponent = GetCurrentComponent(current);
  const [theme, setTheme] = useState<MenuTheme>('dark');

  const changeTheme = (value: boolean) => {
    setTheme(value ? 'dark' : 'light');
  };

  return (
    <Layout className="h-screen " >
      <Sider className="flex-auto " theme={theme} collapsible collapsed={collapsed} onCollapse={(value) => setCollapsed(value)}>
        <Switch
          className="my-2 mx-6"
          checked={theme === 'dark'}
          onChange={changeTheme}
          checkedChildren="Dark"
          unCheckedChildren="Light"
        />
        <br />
        <Menu className="text-left" theme={theme} defaultSelectedKeys={['General']} onClick={onClick} mode="inline" items={items} />
      </Sider>
      <Layout className="h-full w-full">
        {/* <Header className="bg-red-800"/> */}
        <Content className="w-full h-full" >
          <div className="w-full h-full" >
            <CurrentComponent />
          </div>
        </Content>
        {/* <Footer className="bg-purple-500" style={{ textAlign: 'center' }}>
          Ant Design ©{new Date().getFullYear()} Created by Ant UED
        </Footer> */}
      </Layout>
    </Layout>
  );
}

export default Option;
