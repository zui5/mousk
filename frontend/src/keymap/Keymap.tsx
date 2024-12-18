import React, { useEffect, useState } from 'react';

import { KeymapService } from "../../bindings/mousk/service";
interface ShortcutItem {
    key: string;
    shortcut: string;
}

interface ShortcutSection {
    title: string;
    items: ShortcutItem[];
}

const Keymap: React.FC = () => {
    // 定义快捷键数据结构
    const initialShortcutData: ShortcutSection[] = [
        {
            title: 'General Controls',
            items: [
                { key: 'Toggle Mode', shortcut: 'LALT + 0' },
                { key: 'Open Settings', shortcut: 'SPACE + COMMA' },
                { key: 'Reset Settings', shortcut: 'LALT + R' },
                { key: 'Force Quit', shortcut: 'LCONTROL + LSHIFT + A' },
                { key: 'Temporary Quit Control Mode', shortcut: 'Q' },
            ]
        },
        {
            title: 'Mouse Movement',
            items: [
                { key: 'Fast mode', shortcut: 'J, H, L, K' },
                { key: 'Slow mode', shortcut: 'S, A, D, W' },
                { key: 'Speed levels', shortcut: '1, 2, 3, 4, 5' },
            ]
        },
        {
            title: 'Mouse Scroll Simulation',
            items: [
                { key: 'Fast mode', shortcut: 'LSHIFT + J, H, L, K' },
                { key: 'Slow mode', shortcut: 'LSHIFT + S, A, D, W' },
                { key: 'Speed levels', shortcut: 'LShift + 1, 2, 3, 4, 5' },
            ]
        },
        {
            title: 'Mouse Clicks',
            items: [
                { key: 'Left Button Click', shortcut: 'I, Secondary - R' },
                { key: 'Right Button Click', shortcut: 'O, Secondary - T' },
                { key: 'Left Button Hold', shortcut: 'C, Secondary - N' },
            ]
        }
    ];

    // 确保在组件挂载时设置和清除事件监听器
    useEffect(() => {
        // 定义异步函数来调用 a 并获取结果
        const fetchData = async () => {
            keycodeMap= await KeymapService.GetValidKeycodes()
            console.log("keycodeMap init:", keycodeMap)
        };

        fetchData();  // 调用异步函数
        const handleModifierKeyDown = (e: KeyboardEvent) => handleKeyDown(e as unknown as React.KeyboardEvent);
        const handleModifierKeyUp = (e: KeyboardEvent) => handleKeyDown(e as unknown as React.KeyboardEvent);

        window.addEventListener('keydown', handleModifierKeyDown);
        window.addEventListener('keyup', handleModifierKeyUp);

        return () => {
            window.removeEventListener('keydown', handleModifierKeyDown);
            window.removeEventListener('keyup', handleModifierKeyUp);
        };
    }, []);



    const [shortcutData, setShortcutData] = useState<ShortcutSection[]>(initialShortcutData);
    const [editingShortcut, setEditingShortcut] = useState<{ sectionIndex: number, itemIndex: number } | null>(null);

    const handleShortcutClick = (sectionIndex: number, itemIndex: number) => {
        setEditingShortcut({ sectionIndex, itemIndex });
    };

    // 全局数组，用来保存当前按下的修饰键
    const pressedModifiers: string[] = [];
    var keycodeMap: { [key: string]: number } = {}

    const handleKeyDown = async (e: React.KeyboardEvent) => {
        console.log("handleKeyDown", e);
        // 处理修饰键的按下状态
        const handleModifierKey = (code: string, pressed: boolean) => {
            const modifierMap: { [key: string]: string } = {
                'ControlLeft': 'LeftCtrl',
                'ControlRight': 'RightCtrl',
                'ShiftLeft': 'LeftShift',
                'ShiftRight': 'RightShift',
                'AltLeft': 'LeftAlt',
                'AltRight': 'RightAlt',
                'MetaLeft': 'LeftMeta',
                'MetaRight': 'RightMeta'
            };

            const modifierKey = modifierMap[code];

            if (!modifierKey) return;

            const index = pressedModifiers.indexOf(modifierKey);
            if (pressed && index === -1) {
                pressedModifiers.push(modifierKey);
            } else if (!pressed && index !== -1) {
                pressedModifiers.splice(index, 1);
            }
        };



        // 更新修饰键的状态
        if (e.code.includes('Control')) handleModifierKey(e.code, e.ctrlKey);
        if (e.code.includes('Shift')) handleModifierKey(e.code, e.shiftKey);
        if (e.code.includes('Alt')) handleModifierKey(e.code, e.altKey);
        if (e.code.includes('Meta')) handleModifierKey(e.code, e.metaKey);

        if (editingShortcut !== null) {
            const { sectionIndex, itemIndex } = editingShortcut;

            // 检查是否按下的是修饰键
            const isModifierKey = e.key === 'Control' || e.key === 'Shift' || e.key === 'Alt' || e.key === 'Meta';

            // 如果按下的只是修饰键，直接返回
            if (isModifierKey) {
                return;
            }

            // 检查组合键
            const keys: string[] = [...pressedModifiers];
            const key = e.code.replace('Key', '').replace('Digit', '');
            keys.push(key);

            const newShortcut = keys.join('+');

            setShortcutData(prevData => {
                const newData = [...prevData];
                newData[sectionIndex].items[itemIndex].shortcut = newShortcut;
                return newData;
            });

            setEditingShortcut(null);
        }
    };


    return (
        <div className="flex flex-wrap justify-between p-8 rounded-xl w-full h-screen shadow-xl overflow-auto" onKeyDown={handleKeyDown} tabIndex={0}>
            {
                shortcutData.map((section, sectionIndex) => (
                    <div key={sectionIndex} className="flex-1 min-w-[300px] mr-8">
                        <h2 className="text-xl mb-4 border-b-2 border-gray-600 pb-2">{section.title}:</h2>
                        <ul className="pl-5 mb-5">
                            {section.items.map((item, itemIndex) => (
                                <li key={itemIndex} className="mb-3">
                                    {item.key}:
                                    {editingShortcut?.sectionIndex === sectionIndex && editingShortcut.itemIndex === itemIndex ? (
                                        <input
                                            type="text"
                                            value={item.shortcut}
                                            className="bg-gray-700 text-yellow-400 py-1 px-2 rounded-md font-semibold"
                                            readOnly
                                        />
                                    ) : (
                                        <code
                                            className="bg-gray-700 text-yellow-400 py-1 px-2 rounded-md font-semibold cursor-pointer"
                                            onClick={() => handleShortcutClick(sectionIndex, itemIndex)}
                                        >
                                            {item.shortcut}
                                        </code>
                                    )}
                                </li>
                            ))}
                        </ul>
                    </div>
                ))
            }
        </div>
    );
};

export default Keymap;
