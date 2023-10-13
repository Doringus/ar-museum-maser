package com.monkeprojects.armaster;

import com.google.zxing.BarcodeFormat;
import com.google.zxing.WriterException;
import com.google.zxing.client.j2se.MatrixToImageWriter;
import com.google.zxing.common.BitMatrix;
import com.google.zxing.qrcode.QRCodeWriter;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.annotation.Bean;
import org.springframework.core.io.ByteArrayResource;
import org.springframework.core.io.Resource;
import org.springframework.http.HttpHeaders;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.http.converter.BufferedImageHttpMessageConverter;
import org.springframework.http.converter.HttpMessageConverter;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import java.awt.image.BufferedImage;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;

@SpringBootApplication
@RestController
public class ArmasterApplication {

	private static final String MAIN_QR_ID = "507f1f77bcf86cd799439011";
	private static final String MODEL_ID = "00000020f51bb4362eee2a4d";

	@Bean
	public HttpMessageConverter<BufferedImage> createImageHttpMessageConverter() {
		return new BufferedImageHttpMessageConverter();
	}

	public static void main(String[] args) {
		SpringApplication.run(ArmasterApplication.class, args);
	}
	@GetMapping(value = "/main_qr", produces = MediaType.IMAGE_PNG_VALUE)
	public ResponseEntity<BufferedImage> mainQr(@RequestParam(value = "name", defaultValue = "World") String name) throws Exception {
		QRCodeWriter barcodeWriter = new QRCodeWriter();
		BitMatrix bitMatrix =
				barcodeWriter.encode(MAIN_QR_ID, BarcodeFormat.QR_CODE, 200, 200);
		return ResponseEntity.ok(MatrixToImageWriter.toBufferedImage(bitMatrix));
	}

	@GetMapping(value = "/" + MAIN_QR_ID, produces = MediaType.IMAGE_PNG_VALUE)
	public ResponseEntity<BufferedImage> modelQr(@RequestParam(value = "name", defaultValue = "World") String name) throws Exception {
		QRCodeWriter barcodeWriter = new QRCodeWriter();
		BitMatrix bitMatrix =
				barcodeWriter.encode(MODEL_ID, BarcodeFormat.QR_CODE, 200, 200);
		return ResponseEntity.ok(MatrixToImageWriter.toBufferedImage(bitMatrix));
	}

	@GetMapping(value = "/" + MAIN_QR_ID + "/" + MODEL_ID, produces = MediaType.IMAGE_PNG_VALUE)
	public ResponseEntity<Resource> model() throws Exception {
		Path file = Paths.get("D:\\sample.mp3");

		// Get the media type of the file
		String contentType = Files.probeContentType(file);
		if (contentType == null) {
			// Use the default media type
			contentType = MediaType.APPLICATION_OCTET_STREAM_VALUE;
		}

		ByteArrayResource resource = new ByteArrayResource(Files.readAllBytes(file));
		return ResponseEntity.ok()
				.header(HttpHeaders.CONTENT_DISPOSITION, "attachment; filename=\"" + file.getFileName() + "\"")
				.contentLength(Files.size(file))
				.contentType(MediaType.APPLICATION_OCTET_STREAM)
				.body(resource);
	}

}
