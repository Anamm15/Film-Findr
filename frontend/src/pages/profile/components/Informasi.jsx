import { capitalize } from "../../../utils/string";

const Informasi = (props) => {
    const { watch } = props;
    return (
        <div>
            <h2 className="text-2xl font-semibold">{watch.film.judul}</h2>
            <p className="text-gray-600 mt-1">
                <span className="font-medium text-gray-800">Release Date:</span>{" "}
                {watch.film.tanggal_rilis}
            </p>
            <p className="text-gray-600 mt-1">
                <span className="font-medium text-gray-800">Status:</span>{" "}
                {capitalize(watch.film.status)}
            </p>
            <p className="text-gray-600 mt-1">
                <span className="font-medium text-gray-800">Watchlist:</span>{" "}
                {capitalize(watch.status)}
            </p>
        </div>
    )
}

export default Informasi