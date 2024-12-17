import React, { useState } from 'react';

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

    const [shortcutData, setShortcutData] = useState<ShortcutSection[]>(initialShortcutData);
    const [editingShortcut, setEditingShortcut] = useState<{ sectionIndex: number, itemIndex: number } | null>(null);

    const handleShortcutClick = (sectionIndex: number, itemIndex: number) => {
        setEditingShortcut({ sectionIndex, itemIndex });
    };

    const handleKeyDown = (e: React.KeyboardEvent) => {
        if (editingShortcut !== null) {
            const { sectionIndex, itemIndex } = editingShortcut;
            const newShortcut = e.key.toUpperCase(); // 可以根据需求修改此处逻辑

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
