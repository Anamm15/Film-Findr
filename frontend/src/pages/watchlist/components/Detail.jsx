import React, { useState } from "react";
import { WATCH_LIST_STATUS } from "../../../utils/constant";
import Button from "../../../components/Button";
import { updateUserFilm } from "../../../service/userFilm";

const Detail = (props) => {
    const { watchlist } = props;
    const [messages, setMessages] = useState(null);
    const [status, setStatus] = useState();
    const [errorStatus, setErrorStatus] = useState(null);

    const handleUpdate = async () => {
        const data = {
            status: status,
            film_id: watchlist.film.id,
        }

        try {
            const response = await updateUserFilm(watchlist.id, data);
            setMessages(response.message);
            setErrorStatus(false);
        } catch (error) {
            setMessages(error.response.data.error);
            setErrorStatus(true);
        }
    };

    return (
        <div className="mt-4 flex flex-col gap-2">
            <div className="flex flex-col md:flex-row gap-2 items-center">
                <div className="relative max-w-sm text-lg w-full md:w-auto">
                    <select
                        className="appearance-none px-4 py-1 pr-10 rounded-lg border border-gray-300 shadow-sm focus:outline-none focus:ring-2 focus:ring-blue-500 transition duration-200 w-full md:w-auto"
                        onChange={(e) => setStatus(e.target.value)}
                    >
                        <option value="">Select Status</option>
                        {WATCH_LIST_STATUS.map((status) => (
                            <option key={status} value={status}>
                                {status}
                            </option>
                        ))}
                    </select>

                    {/* Custom arrow */}
                    <div className="pointer-events-none absolute inset-y-0 right-4 flex items-center text-gray-500">
                        <svg
                            className="w-5 h-5"
                            fill="none"
                            stroke="currentColor"
                            viewBox="0 0 24 24"
                        >
                            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M19 9l-7 7-7-7" />
                        </svg>
                    </div>
                </div>
                <Button
                    type="button"
                    className="rounded-lg py-[7px] w-full md:w-auto"
                    onClick={handleUpdate}
                >
                    Change Status
                </Button>
            </div>
            {messages && (
                <p className={`${errorStatus ? "text-red-600" : "text-green-600"} text-sm font-medium`}>
                    {messages}
                </p>
            )}
        </div>
    )
}

export default Detail;