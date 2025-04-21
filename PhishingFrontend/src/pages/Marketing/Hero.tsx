import { Display } from "@fluentui/react-components";

const Hero = () => {
  return (
    <>
      <div
        style={{
          position: "absolute",
          top: 0,
          left: 0,
          zIndex: -10,
          height: "100vh",
          width: "100vw",
          backgroundColor: 'oklch(0.985 0.002 247.839)',
        }}
      >
        <div
          style={{
            position: "absolute",
            bottom: "auto",
            left: "auto",
            right: 0,
            top: 0,
            height: "500px",
            width: "500px",
            transform: "translate(-70%, 25%)",
            borderRadius: "9999px", // full rounding
            backgroundColor: "rgba(0, 120, 212, 0.5)", // "rgba(173, 109, 244, 0.5)",
            opacity: 0.5,
            filter: "blur(80px)",
          }}
        ></div>
      </div>
      <div
        style={{
          marginTop: "5rem",
          maxWidth: "70rem",
          textAlign: "center",
          margin: "auto"
        }}
      >
        <Display
          style={{
            lineHeight: 1.1
          }}
        >
          Gemeinsam gegen <strong>Phishing und Cyberbetrug</strong> für eine <strong>sichere digitale Zukunft</strong>.
        </Display>
      </div>
    </>
  );
}

export default Hero;