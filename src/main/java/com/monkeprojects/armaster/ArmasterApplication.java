package com.monkeprojects.armaster;

import com.google.zxing.BarcodeFormat;
import com.google.zxing.client.j2se.MatrixToImageWriter;
import com.google.zxing.common.BitMatrix;
import com.google.zxing.qrcode.QRCodeWriter;
import com.monkeprojects.armaster.domain.Exhibition;
import com.monkeprojects.armaster.domain.File;
import com.monkeprojects.armaster.repository.ExhibitionRepository;
import com.monkeprojects.armaster.repository.FileRepository;
import org.springframework.beans.factory.annotation.Autowired;
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

	// Used only for first demo
	String exhibitionId;
	String modelId ="test";

	@Autowired
	private FileRepository fileRepository;

	@Autowired
	private ExhibitionRepository exhibitionRepository;

	@Bean
	public HttpMessageConverter<BufferedImage> createImageHttpMessageConverter() {
		return new BufferedImageHttpMessageConverter();
	}

	public static void main(String[] args) {
		SpringApplication.run(ArmasterApplication.class, args);
	}

	/**
	 * Used only for first demo
	 * Remove later
	 */
	@GetMapping(value = "/mock_post")
	public void mockPost(@RequestParam(name = "path") String pathToModel) throws Exception {
		if(pathToModel == null) {
			throw new Exception();
		}
	//	"/home/dskom/home_projects/Archive.zip"
		var exhibition = exhibitionRepository.insert(new Exhibition("Test exhibition"));
		this.exhibitionId = exhibition.getId();
		Path path = Paths.get(pathToModel);
		var model = fileRepository.insert(new File(path.getFileName().toString(), Files.size(path), Files.readAllBytes(path)));
		modelId = model.getId();
	}

	@GetMapping(value = "/data")
	public ResponseEntity<Resource> localModel(@RequestParam(name = "model_id") String modelId) throws Exception {
		if(modelId == null) {
			throw new Exception();
		}

		var model = fileRepository.findById(modelId);
		if(model.isEmpty()) {
			throw new Exception();
		}
		ByteArrayResource resource = new ByteArrayResource(model.get().getContent());
		return ResponseEntity.ok()
				.header(HttpHeaders.CONTENT_DISPOSITION, "attachment; filename=\"" + model.get().getName() + "\"")
				.contentLength(model.get().getSize())
				.contentType(MediaType.APPLICATION_OCTET_STREAM)
				.body(resource);
	}

	@GetMapping(value = "/main_qr", produces = MediaType.IMAGE_PNG_VALUE)
	public ResponseEntity<BufferedImage> mainQr() throws Exception {
		QRCodeWriter barcodeWriter = new QRCodeWriter();
		BitMatrix bitMatrix =
				barcodeWriter.encode(modelId, BarcodeFormat.QR_CODE, 200, 200);
		return ResponseEntity.ok(MatrixToImageWriter.toBufferedImage(bitMatrix));
	}

}
