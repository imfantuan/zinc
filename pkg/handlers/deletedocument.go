package handlers

import (
	"net/http"

	"github.com/blugelabs/bluge"
	"github.com/gin-gonic/gin"
	"github.com/prabhatsharma/zinc/pkg/core"
)

func DeleteDocument(c *gin.Context) {
	indexName := c.Param("target")
	query_id := c.Param("id")

	index, exists := core.GetIndex(indexName)
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "index does not exists"})
		return
	}

	bdoc := bluge.NewDocument(query_id)
	bdoc.AddField(bluge.NewCompositeFieldExcluding("_all", nil))
	err := index.Writer.Delete(bdoc.ID())
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Deleted", "index": indexName, "id": query_id})
	}
}
