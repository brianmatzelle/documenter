"use client";

import { createContext, useContext, useState } from "react";

interface DocContextType {
	doc: string;
	setDoc: (doc: string) => void;
}

export const DocContext = createContext<DocContextType>({
	doc: "",
	setDoc: () => {},
});

export const useDocContext = () => useContext(DocContext);

export const DocProvider = ({ children }: { children: React.ReactNode }) => {
	const [doc, setDoc] = useState("");

	return <DocContext.Provider value={{ doc, setDoc }}>{children}</DocContext.Provider>;
};
