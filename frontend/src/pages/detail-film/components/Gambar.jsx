import { useState } from "react";
import { ChevronLeft, ChevronRight, Image as ImageIcon } from "lucide-react";

const Gambar = ({ film }) => {
    const [currentIndex, setCurrentIndex] = useState(0);
    const images = film?.film_gambar || [];

    if (images.length === 0) {
        return (
            <div className="h-[450px] w-[300px] bg-gray-200 rounded-lg flex flex-col items-center justify-center text-gray-400 mx-auto md:mx-0">
                <ImageIcon size={48} />
                <span className="text-sm mt-2">No Image</span>
            </div>
        );
    }

    const prevSlide = () => {
        const isFirstSlide = currentIndex === 0;
        const newIndex = isFirstSlide ? images.length - 1 : currentIndex - 1;
        setCurrentIndex(newIndex);
    };

    const nextSlide = () => {
        const isLastSlide = currentIndex === images.length - 1;
        const newIndex = isLastSlide ? 0 : currentIndex + 1;
        setCurrentIndex(newIndex);
    };

    return (
        <div className="relative h-[450px] w-[300px] mx-auto md:mx-0 group">
            {/* Main Image */}
            <div className="w-full h-full rounded-xl overflow-hidden shadow-lg bg-gray-900">
                <img
                    src={images[currentIndex].url}
                    alt={`Scene ${currentIndex + 1}`}
                    className="w-full h-full object-cover transition-opacity duration-500"
                />
            </div>

            {/* Navigation Arrows - Muncul saat hover (opsional) atau selalu ada */}
            {images.length > 1 && (
                <>
                    {/* Left Arrow */}
                    <button
                        onClick={prevSlide}
                        className="absolute top-1/2 left-2 -translate-y-1/2 p-2 rounded-full bg-black/30 hover:bg-black/60 text-white transition-all backdrop-blur-sm border border-white/10"
                    >
                        <ChevronLeft size={24} />
                    </button>

                    {/* Right Arrow */}
                    <button
                        onClick={nextSlide}
                        className="absolute top-1/2 right-2 -translate-y-1/2 p-2 rounded-full bg-black/30 hover:bg-black/60 text-white transition-all backdrop-blur-sm border border-white/10"
                    >
                        <ChevronRight size={24} />
                    </button>

                    {/* Slide Indicator (1/5) */}
                    <div className="absolute bottom-4 right-4 bg-black/60 backdrop-blur-md text-white text-xs font-medium px-3 py-1 rounded-full border border-white/10">
                        {currentIndex + 1} / {images.length}
                    </div>

                    {/* Opsional: Dot Indicator di tengah bawah */}
                    <div className="absolute bottom-4 left-1/2 -translate-x-1/2 flex gap-1.5">
                        {images.map((_, slideIndex) => (
                            <div
                                key={slideIndex}
                                onClick={() => setCurrentIndex(slideIndex)}
                                className={`w-2 h-2 rounded-full cursor-pointer transition-all ${currentIndex === slideIndex ? "bg-white w-4" : "bg-white/50 hover:bg-white/80"
                                    }`}
                            />
                        ))}
                    </div>
                </>
            )}
        </div>
    );
};

export default Gambar;