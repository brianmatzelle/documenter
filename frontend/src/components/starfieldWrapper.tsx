"use client";
import { Starfield } from "battlezone-shapes"

export default function StarfieldWrapper() {
	return <div className="fixed inset-0 -z-10">
		<Starfield />
	</div>
}
