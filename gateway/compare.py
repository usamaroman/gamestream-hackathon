import cv2

img1 = "img1.png"
img2 = "img2.png"

def difference(b1: list, b2: list):
    # Open in "wb" mode to
    # write a new file, or
    # "ab" mode to append
    with open(img1, "wb") as binary_file:
    	# Write bytes to file
    	binary_file.write(b1)

    with open(img2, "wb") as binary_file:
        # Write bytes to file
        binary_file.write(b2)

    # load images
    image1 = cv2.imread(img1)
    image2 = cv2.imread(img2)

    # compute difference
    difference = cv2.subtract(image1, image2)

    # color the mask red
    Conv_hsv_Gray = cv2.cvtColor(difference, cv2.COLOR_BGR2GRAY)
    ret, mask = cv2.threshold(Conv_hsv_Gray, 0, 255,cv2.THRESH_BINARY_INV |cv2.THRESH_OTSU)
    difference[mask != 255] = [0, 0, 255]

    # add the red mask to the images to make the differences obvious
    image1[mask != 255] = [0, 0, 255]
    image2[mask != 255] = [0, 0, 255]

    # store images
    cv2.imwrite('diffOverImage1.png', image1)
    cv2.imwrite('diffOverImage2.png', image1)
    cv2.imwrite('diff.png', difference)