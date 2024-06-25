# This R script takes a distance matrix between genomics samples as input.
# It produces two plots. First, a heat map of the distances in the samples. Second, a PCoA plot
# colored by the location.

# Import needed libraries. Please install each of these if needed by visiting "Tools" --> "Package Installer".

library(ggplot2)
library(ggcorrplot)
library(reshape)
library(stringr)
library(ape) #ape library to compute PCoA of our matrix

# Now set working directory, the location where RStudio should look for data.

# To do so, select "Session" --> "Set Working Directory" --> "To Source File Location". This will ensure that RStudio knows to look in the directory where this file is contained for data.

#Now, we need to set the year. Set this equal to 2019 or 2023 as the case may be 

year <- "2023"

# Part 1: generate box plots of evenness by location using Simpson's index.

# first, read in the table

simpsonTable <- read.csv(file=paste("Matrices/SimpsonMatrix_", year, ".csv",sep=""))

simpsonColumns <- data.frame(simpsonTable)

# Add column to our data frame to represent the location
Location <- sub("\\_.*", "", simpsonColumns$Sample) # parse out just the location from sample name
cbind(simpsonColumns, Location) # adding column

# Now we plot box plots where we have evenness lumped by location.
ggplot(simpsonColumns, aes(x=Location, y=SimpsonsIndex, fill=Location)) + geom_boxplot() +ggtitle("Simpson's Index over location")
ggsave(paste("Plots/SimpsonsBoxPlots_",year,".png", sep=""), width=8, height=6, dpi=1000)


# Part 2: generate a heat map of the distance values

# Read in the file and process the table.
table <- read.csv(file=paste("Matrices/BrayCurtisBetaDiversityMatrix_",year,".csv",sep=""))

# trim the first column out because it only contains names 
table <- table[-c(1)]
table <- table[-c(ncol(table))]


matrix <- as.matrix(table)

# the following code is just all the technical stuff to build a heatmap out of the distance matrix.
co=melt(matrix)
head(co)
ggplot(co, aes(X1, X2)) + # x and y axes => Var1 and Var2
geom_tile(aes(fill = value)) + # background colours are mapped according to the value column
scale_fill_gradient2(low = "#6D9EC1", 
                     mid = "white", 
                     high = "#E46726", 
                     midpoint = 0.5, limit= c(0,1.0)) + 
theme(panel.grid.major.x=element_blank(), #no gridlines
      panel.grid.minor.x=element_blank(), 
      panel.grid.major.y=element_blank(), 
      panel.grid.minor.y=element_blank(),
      panel.background=element_rect(fill="white"), # background=white
      axis.text.x = element_text(angle=90, hjust = 1,vjust=1,size = 8,face = "bold"),
      plot.title = element_text(size=20,face="bold"),
      axis.text.y = element_text(size = 8,face = "bold")) + 
ggtitle("Bray Curtis Heat Map") + 
theme(legend.title=element_text(face="bold", size=14)) + 
scale_x_discrete(name="") +
scale_y_discrete(name="")
ggsave(paste("Plots/BrayCurtisHeatMap_",year,".png", sep=""), width=8, height=6, dpi=1000)

# Let's do all this again for Jaccard 
# Read in the file and process the table.
table <- read.csv(file=paste("Matrices/JaccardBetaDiversityMatrix_", year, ".csv", sep=""))

# trim the first column out because it only contains names 
table <- table[-c(1)]
table <- table[-c(ncol(table))]


matrix <- as.matrix(table)

# the following code is just all the technical stuff to build a heatmap out of the distance matrix.
co=melt(matrix)
head(co)
ggplot(co, aes(X1, X2)) + # x and y axes => Var1 and Var2
   geom_tile(aes(fill = value)) + # background colours are mapped according to the value column
   scale_fill_gradient2(low = "#6D9EC1", 
                        mid = "white", 
                        high = "#E46726", 
                        midpoint = 0.5, limit= c(0,1.0)) + 
   theme(panel.grid.major.x=element_blank(), #no gridlines
         panel.grid.minor.x=element_blank(), 
         panel.grid.major.y=element_blank(), 
         panel.grid.minor.y=element_blank(),
         panel.background=element_rect(fill="white"), # background=white
         axis.text.x = element_text(angle=90, hjust = 1,vjust=1,size = 8,face = "bold"),
         plot.title = element_text(size=20,face="bold"),
         axis.text.y = element_text(size = 8,face = "bold")) + 
   ggtitle("Jaccard Heat Map") + 
   theme(legend.title=element_text(face="bold", size=14)) + 
   scale_x_discrete(name="") +
   scale_y_discrete(name="")
ggsave(paste("Plots/JaccardHeatMap_",year,".png", sep=""), width=8, height=6, dpi=1000)

# Part 3: generate a PCoA plot of the data and color-code by location.

# First, Bray Curtis
# Read in the file and process the table.
table2 <- read.csv(file=paste("Matrices/BrayCurtisBetaDiversityMatrix_", year, ".csv", sep=""))

#trim the first column out because it only contains names 
table2 <- table2[-c(1)]
table2 <- table2[-c(ncol(table2))] # trim out weird extra column at end of the matrix file

matrix2 <- as.matrix(table2)

pcoa_data <- pcoa(matrix2, correction="none", rn=NULL)
pcoa_vectors <- data.frame(pcoa_data$vectors)
# columns contains a vector for each point after PCoA tries to assign data points to vectors to preserve distances between points.

colnames(table2)

# Add column to our data frame to represent the location
Location2 <- sub("\\_.*", "", colnames(table2)) # parse out just the location from sample name
cbind(pcoa_vectors, Location2) # adding column

# Now, plot the data, colored by location.
ggplot(pcoa_vectors, aes(x=Axis.1, y=Axis.2, color=Location2)) + geom_point()
ggsave(paste("Plots/BrayCurtisPCoA_",year,".png", sep=""), width=8, height=6, dpi=1000)


# Second, Jaccard 

# Read in the file and process the table.
table2 <- read.csv(file=paste("Matrices/JaccardBetaDiversityMatrix_", year, ".csv", sep=""))

#trim the first column out because it only contains names 
table2 <- table2[-c(1)]
table2 <- table2[-c(ncol(table2))] # trim out weird extra column at end of the matrix file

matrix2 <- as.matrix(table2)

pcoa_data <- pcoa(matrix2, correction="none", rn=NULL)
pcoa_vectors <- data.frame(pcoa_data$vectors)
# columns contains a vector for each point after PCoA tries to assign data points to vectors to preserve distances between points.

colnames(table2)

# Add column to our data frame to represent the location
Location2 <- sub("\\_.*", "", colnames(table2)) # parse out just the location from sample name
cbind(pcoa_vectors, Location2) # adding column

# Now, plot the data, colored by location.
ggplot(pcoa_vectors, aes(x=Axis.1, y=Axis.2, color=Location2)) + geom_point()


ggsave(paste("Plots/JaccardPCoA_",year,".png", sep=""), width=8, height=6, dpi=1000)

