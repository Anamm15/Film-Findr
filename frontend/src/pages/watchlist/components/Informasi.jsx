import { useMediaQuery } from "react-responsive";
import { capitalize } from "../../../utils/string";

const Informasi = (props) => {
    const { watch } = props;
    const isMobile = useMediaQuery({ maxWidth: 640 });

    function sliceTitle(title, length) {
        if (title.length > length) {
            return title.slice(0, length) + "...";
        }
        return title;
    }

    return (
        <div>
            <h2 className="text-xl md:text-2xl font-semibold text-text">
                {isMobile ? sliceTitle(watch.film.judul, 28) : watch.film.judul}
            </h2>
            <p className="text-secondary mt-1">
                <span className="text-sm md:text-md font-medium text-gray-800">Release Date:</span>{" "}
                {watch.film.tanggal_rilis}
            </p>
            <p className="text-secondary mt-1">
                <span className="text-sm md:text-md font-medium text-gray-800">Status:</span>{" "}
                {capitalize(watch.film.status)}
            </p>
            <p className="text-secondary mt-1">
                <span className="text-sm md:text-md font-medium text-gray-800">Watchlist:</span>{" "}
                {capitalize(watch.status)}
            </p>
        </div>
    )
}

export default Informasi