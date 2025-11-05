package article

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"tutorial/api/initializers"

	"github.com/Azure/azure-sdk-for-go/sdk/data/azcosmos"
	"github.com/gin-gonic/gin"
)

func AddArticle(c *gin.Context) {
	var article Article
	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	partitionKey := azcosmos.NewPartitionKeyString(article.ID)

	ctx := context.TODO()

	bytes, err := json.Marshal(article)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	response, err := initializers.Container.UpsertItem(ctx, partitionKey, bytes, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"id":       article.ID,
		"activity": response.ActivityID,
		"charge":   response.RequestCharge,
	})

}

func CheckAnItem(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id richiesto nel query param"})
		return
	}

	partitionKey := azcosmos.NewPartitionKeyString(id)

	context := context.TODO()

	response, err := initializers.Container.ReadItem(context, partitionKey, id, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "id": id})
		return
	}
	var article Article

	if response.RawResponse.StatusCode == 200 {
		err := json.Unmarshal(response.Value, &article)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, article)
	}
}

func ListNArticles(c *gin.Context) {
	qty := c.Query("qty")
	intQty, err := strconv.ParseInt(qty, 10, 64)
	if err != nil {
		log.Fatal("failed to convert string to integer", err)
	}

	partitionKey := azcosmos.NewPartitionKey()

	query := "SELECT * FROM c"

	pager := initializers.Container.NewQueryItemsPager(query, partitionKey, nil)

	items := []Article{}
	count := 0

	for pager.More() && count < int(intQty) {
		response, err := pager.NextPage(context.TODO())
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"errore": err.Error()})
			return
		}
		for _, bytes := range response.Items {
			if count >= int(intQty) {
				break
			}
			item := Article{}
			err := json.Unmarshal(bytes, &item)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"errore": err.Error()})
				return
			}
			items = append(items, item)
			count++
		}
	}
	c.JSON(http.StatusOK, items)
}
